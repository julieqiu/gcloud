package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	netapp "cloud.google.com/go/netapp/apiv1"
	"cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "netapp",
				Usage: "manage NetApp API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "active-directories",
						Usage: "Manage active-directories resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list active-directories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListActiveDirectoriesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListActiveDirectories(ctx, req)
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
								Usage: "describe active-directories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "active_directory", Usage: "The active_directory.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", cmd.String("project"), cmd.String("location"), cmd.String("active_directory"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetActiveDirectoryRequest{Name: name}
									resp, err := client.GetActiveDirectory(ctx, req)
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
								Usage: "create active-directories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "active-directory-id", Usage: "The active directory id.", Required: true},
									&cli.StringFlag{Name: "domain", Usage: "The domain.", Required: true},
									&cli.StringFlag{Name: "site", Usage: "The site.", Required: false},
									&cli.StringFlag{Name: "dns", Usage: "The dns.", Required: true},
									&cli.StringFlag{Name: "net-bios-prefix", Usage: "The net bios prefix.", Required: true},
									&cli.StringFlag{Name: "organizational-unit", Usage: "The organizational unit.", Required: false},
									&cli.BoolFlag{Name: "aes-encryption", Usage: "The aes encryption.", Required: false},
									&cli.StringFlag{Name: "username", Usage: "The username.", Required: true},
									&cli.StringFlag{Name: "password", Usage: "The password.", Required: true},
									&cli.StringFlag{Name: "kdc-hostname", Usage: "The kdc hostname.", Required: false},
									&cli.StringFlag{Name: "kdc-ip", Usage: "The kdc ip.", Required: false},
									&cli.BoolFlag{Name: "nfs-users-with-ldap", Usage: "The nfs users with ldap.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "ldap-signing", Usage: "The ldap signing.", Required: false},
									&cli.BoolFlag{Name: "encrypt-dc-connections", Usage: "The encrypt dc connections.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateActiveDirectoryRequest{Parent: parent}
									req.ActiveDirectoryId = cmd.String("active-directory-id")
									req.ActiveDirectory = &netapppb.ActiveDirectory{
										Domain:               cmd.String("domain"),
										Site:                 cmd.String("site"),
										Dns:                  cmd.String("dns"),
										NetBiosPrefix:        cmd.String("net-bios-prefix"),
										OrganizationalUnit:   cmd.String("organizational-unit"),
										AesEncryption:        cmd.Bool("aes-encryption"),
										Username:             cmd.String("username"),
										Password:             cmd.String("password"),
										KdcHostname:          cmd.String("kdc-hostname"),
										KdcIp:                cmd.String("kdc-ip"),
										NfsUsersWithLdap:     cmd.Bool("nfs-users-with-ldap"),
										Description:          cmd.String("description"),
										LdapSigning:          cmd.Bool("ldap-signing"),
										EncryptDcConnections: cmd.Bool("encrypt-dc-connections"),
									}
									op, err := client.CreateActiveDirectory(ctx, req)
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
								Usage: "update active-directories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "active_directory", Usage: "The active_directory.", Required: true},
									&cli.StringFlag{Name: "domain", Usage: "The domain.", Required: false},
									&cli.StringFlag{Name: "site", Usage: "The site.", Required: false},
									&cli.StringFlag{Name: "dns", Usage: "The dns.", Required: false},
									&cli.StringFlag{Name: "net-bios-prefix", Usage: "The net bios prefix.", Required: false},
									&cli.StringFlag{Name: "organizational-unit", Usage: "The organizational unit.", Required: false},
									&cli.BoolFlag{Name: "aes-encryption", Usage: "The aes encryption.", Required: false},
									&cli.StringFlag{Name: "username", Usage: "The username.", Required: false},
									&cli.StringFlag{Name: "password", Usage: "The password.", Required: false},
									&cli.StringFlag{Name: "kdc-hostname", Usage: "The kdc hostname.", Required: false},
									&cli.StringFlag{Name: "kdc-ip", Usage: "The kdc ip.", Required: false},
									&cli.BoolFlag{Name: "nfs-users-with-ldap", Usage: "The nfs users with ldap.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "ldap-signing", Usage: "The ldap signing.", Required: false},
									&cli.BoolFlag{Name: "encrypt-dc-connections", Usage: "The encrypt dc connections.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", cmd.String("project"), cmd.String("location"), cmd.String("active_directory"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateActiveDirectoryRequest{}
									req.ActiveDirectory = &netapppb.ActiveDirectory{
										Name:                 name,
										Domain:               cmd.String("domain"),
										Site:                 cmd.String("site"),
										Dns:                  cmd.String("dns"),
										NetBiosPrefix:        cmd.String("net-bios-prefix"),
										OrganizationalUnit:   cmd.String("organizational-unit"),
										AesEncryption:        cmd.Bool("aes-encryption"),
										Username:             cmd.String("username"),
										Password:             cmd.String("password"),
										KdcHostname:          cmd.String("kdc-hostname"),
										KdcIp:                cmd.String("kdc-ip"),
										NfsUsersWithLdap:     cmd.Bool("nfs-users-with-ldap"),
										Description:          cmd.String("description"),
										LdapSigning:          cmd.Bool("ldap-signing"),
										EncryptDcConnections: cmd.Bool("encrypt-dc-connections"),
									}
									var paths []string
									if cmd.IsSet("domain") {
										paths = append(paths, "domain")
									}
									if cmd.IsSet("site") {
										paths = append(paths, "site")
									}
									if cmd.IsSet("dns") {
										paths = append(paths, "dns")
									}
									if cmd.IsSet("net-bios-prefix") {
										paths = append(paths, "net_bios_prefix")
									}
									if cmd.IsSet("organizational-unit") {
										paths = append(paths, "organizational_unit")
									}
									if cmd.IsSet("aes-encryption") {
										paths = append(paths, "aes_encryption")
									}
									if cmd.IsSet("username") {
										paths = append(paths, "username")
									}
									if cmd.IsSet("password") {
										paths = append(paths, "password")
									}
									if cmd.IsSet("kdc-hostname") {
										paths = append(paths, "kdc_hostname")
									}
									if cmd.IsSet("kdc-ip") {
										paths = append(paths, "kdc_ip")
									}
									if cmd.IsSet("nfs-users-with-ldap") {
										paths = append(paths, "nfs_users_with_ldap")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("ldap-signing") {
										paths = append(paths, "ldap_signing")
									}
									if cmd.IsSet("encrypt-dc-connections") {
										paths = append(paths, "encrypt_dc_connections")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateActiveDirectory(ctx, req)
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
								Usage: "delete active-directories",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "active_directory", Usage: "The active_directory.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/activeDirectories/%s", cmd.String("project"), cmd.String("location"), cmd.String("active_directory"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteActiveDirectoryRequest{Name: name}
									op, err := client.DeleteActiveDirectory(ctx, req)
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
						Name:  "backup-policies",
						Usage: "Manage backup-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create backup-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup-policy-id", Usage: "The backup policy id.", Required: true},
									&cli.IntFlag{Name: "daily-backup-limit", Usage: "The daily backup limit.", Required: false},
									&cli.IntFlag{Name: "weekly-backup-limit", Usage: "The weekly backup limit.", Required: false},
									&cli.IntFlag{Name: "monthly-backup-limit", Usage: "The monthly backup limit.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "enabled", Usage: "The enabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateBackupPolicyRequest{Parent: parent}
									req.BackupPolicyId = cmd.String("backup-policy-id")
									req.BackupPolicy = &netapppb.BackupPolicy{
										DailyBackupLimit:   int32(cmd.Int("daily-backup-limit")),
										WeeklyBackupLimit:  int32(cmd.Int("weekly-backup-limit")),
										MonthlyBackupLimit: int32(cmd.Int("monthly-backup-limit")),
										Description:        cmd.String("description"),
										Enabled:            cmd.Bool("enabled"),
									}
									op, err := client.CreateBackupPolicy(ctx, req)
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
								Name:  "describe",
								Usage: "describe backup-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_policy", Usage: "The backup_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_policy"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetBackupPolicyRequest{Name: name}
									resp, err := client.GetBackupPolicy(ctx, req)
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
								Usage: "list backup-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListBackupPoliciesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListBackupPolicies(ctx, req)
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
								Usage: "update backup-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_policy", Usage: "The backup_policy.", Required: true},
									&cli.IntFlag{Name: "daily-backup-limit", Usage: "The daily backup limit.", Required: false},
									&cli.IntFlag{Name: "weekly-backup-limit", Usage: "The weekly backup limit.", Required: false},
									&cli.IntFlag{Name: "monthly-backup-limit", Usage: "The monthly backup limit.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "enabled", Usage: "The enabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_policy"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateBackupPolicyRequest{}
									req.BackupPolicy = &netapppb.BackupPolicy{
										Name:               name,
										DailyBackupLimit:   int32(cmd.Int("daily-backup-limit")),
										WeeklyBackupLimit:  int32(cmd.Int("weekly-backup-limit")),
										MonthlyBackupLimit: int32(cmd.Int("monthly-backup-limit")),
										Description:        cmd.String("description"),
										Enabled:            cmd.Bool("enabled"),
									}
									var paths []string
									if cmd.IsSet("daily-backup-limit") {
										paths = append(paths, "daily_backup_limit")
									}
									if cmd.IsSet("weekly-backup-limit") {
										paths = append(paths, "weekly_backup_limit")
									}
									if cmd.IsSet("monthly-backup-limit") {
										paths = append(paths, "monthly_backup_limit")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("enabled") {
										paths = append(paths, "enabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateBackupPolicy(ctx, req)
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
								Usage: "delete backup-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_policy", Usage: "The backup_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_policy"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteBackupPolicyRequest{Name: name}
									op, err := client.DeleteBackupPolicy(ctx, req)
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
						Name:  "backup-vaults",
						Usage: "Manage backup-vaults resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create backup-vaults",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup-vault-id", Usage: "The backup vault id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "backup-region", Usage: "The backup region.", Required: false},
									&cli.StringFlag{Name: "kms-config", Usage: "The kms config.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateBackupVaultRequest{Parent: parent}
									req.BackupVaultId = cmd.String("backup-vault-id")
									req.BackupVault = &netapppb.BackupVault{
										Description:  cmd.String("description"),
										BackupRegion: cmd.String("backup-region"),
										KmsConfig:    cmd.String("kms-config"),
									}
									op, err := client.CreateBackupVault(ctx, req)
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
								Name:  "describe",
								Usage: "describe backup-vaults",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetBackupVaultRequest{Name: name}
									resp, err := client.GetBackupVault(ctx, req)
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
								Usage: "list backup-vaults",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListBackupVaultsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListBackupVaults(ctx, req)
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
								Usage: "update backup-vaults",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "backup-region", Usage: "The backup region.", Required: false},
									&cli.StringFlag{Name: "kms-config", Usage: "The kms config.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateBackupVaultRequest{}
									req.BackupVault = &netapppb.BackupVault{
										Name:         name,
										Description:  cmd.String("description"),
										BackupRegion: cmd.String("backup-region"),
										KmsConfig:    cmd.String("kms-config"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("backup-region") {
										paths = append(paths, "backup_region")
									}
									if cmd.IsSet("kms-config") {
										paths = append(paths, "kms_config")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateBackupVault(ctx, req)
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
								Usage: "delete backup-vaults",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteBackupVaultRequest{Name: name}
									op, err := client.DeleteBackupVault(ctx, req)
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
						Name:  "backups",
						Usage: "Manage backups resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
									&cli.StringFlag{Name: "backup-id", Usage: "The backup id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "source-volume", Usage: "The source volume.", Required: false},
									&cli.StringFlag{Name: "source-snapshot", Usage: "The source snapshot.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateBackupRequest{Parent: parent}
									req.BackupId = cmd.String("backup-id")
									req.Backup = &netapppb.Backup{
										Description:    cmd.String("description"),
										SourceVolume:   cmd.String("source-volume"),
										SourceSnapshot: cmd.String("source-snapshot"),
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
								Name:  "describe",
								Usage: "describe backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"), cmd.String("backup"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetBackupRequest{Name: name}
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
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListBackupsRequest{Parent: parent}
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
								Name:  "delete",
								Usage: "delete backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"), cmd.String("backup"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteBackupRequest{Name: name}
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
								Name:  "update",
								Usage: "update backups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backup_vault", Usage: "The backup_vault.", Required: true},
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "source-volume", Usage: "The source volume.", Required: false},
									&cli.StringFlag{Name: "source-snapshot", Usage: "The source snapshot.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s/backups/%s", cmd.String("project"), cmd.String("location"), cmd.String("backup_vault"), cmd.String("backup"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateBackupRequest{}
									req.Backup = &netapppb.Backup{
										Name:           name,
										Description:    cmd.String("description"),
										SourceVolume:   cmd.String("source-volume"),
										SourceSnapshot: cmd.String("source-snapshot"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("source-volume") {
										paths = append(paths, "source_volume")
									}
									if cmd.IsSet("source-snapshot") {
										paths = append(paths, "source_snapshot")
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
						},
					},
					{
						Name:  "host-groups",
						Usage: "Manage host-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list host-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListHostGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListHostGroups(ctx, req)
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
								Usage: "describe host-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "host_group", Usage: "The host_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("host_group"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetHostGroupRequest{Name: name}
									resp, err := client.GetHostGroup(ctx, req)
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
								Usage: "create host-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "host-group-id", Usage: "The host group id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateHostGroupRequest{Parent: parent}
									req.HostGroupId = cmd.String("host-group-id")
									req.HostGroup = &netapppb.HostGroup{
										Description: cmd.String("description"),
									}
									op, err := client.CreateHostGroup(ctx, req)
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
								Usage: "update host-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "host_group", Usage: "The host_group.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("host_group"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateHostGroupRequest{}
									req.HostGroup = &netapppb.HostGroup{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateHostGroup(ctx, req)
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
								Usage: "delete host-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "host_group", Usage: "The host_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/hostGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("host_group"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteHostGroupRequest{Name: name}
									op, err := client.DeleteHostGroup(ctx, req)
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
						Name:  "kms-configs",
						Usage: "Manage kms-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list kms-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListKmsConfigsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListKmsConfigs(ctx, req)
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
								Usage: "create kms-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "kms-config-id", Usage: "The kms config id.", Required: true},
									&cli.StringFlag{Name: "crypto-key-name", Usage: "The crypto key name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateKmsConfigRequest{Parent: parent}
									req.KmsConfigId = cmd.String("kms-config-id")
									req.KmsConfig = &netapppb.KmsConfig{
										CryptoKeyName: cmd.String("crypto-key-name"),
										Description:   cmd.String("description"),
									}
									op, err := client.CreateKmsConfig(ctx, req)
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
								Name:  "describe",
								Usage: "describe kms-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "kms_config", Usage: "The kms_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms_config"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetKmsConfigRequest{Name: name}
									resp, err := client.GetKmsConfig(ctx, req)
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
								Usage: "update kms-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "kms_config", Usage: "The kms_config.", Required: true},
									&cli.StringFlag{Name: "crypto-key-name", Usage: "The crypto key name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms_config"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateKmsConfigRequest{}
									req.KmsConfig = &netapppb.KmsConfig{
										Name:          name,
										CryptoKeyName: cmd.String("crypto-key-name"),
										Description:   cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("crypto-key-name") {
										paths = append(paths, "crypto_key_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateKmsConfig(ctx, req)
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
								Name:  "encrypt",
								Usage: "encrypt kms-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "kms_config", Usage: "The kms_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms_config"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.EncryptVolumesRequest{Name: name}
									op, err := client.EncryptVolumes(ctx, req)
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
								Name:  "verify",
								Usage: "verify kms-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "kms_config", Usage: "The kms_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms_config"))
									fmt.Printf("Executing verify on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete kms-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "kms_config", Usage: "The kms_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/kmsConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("kms_config"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteKmsConfigRequest{Name: name}
									op, err := client.DeleteKmsConfig(ctx, req)
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
									client, err := netapp.NewClient(ctx)
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
									client, err := netapp.NewClient(ctx)
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
						Name:  "ontap",
						Usage: "Manage ontap resources",
						Commands: []*cli.Command{
							{
								Name:  "execute-ontap-post",
								Usage: "execute-ontap-post ontap",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing execute-ontap-post...")
									return nil
								},
							},
							{
								Name:  "execute-ontap-get",
								Usage: "execute-ontap-get ontap",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing execute-ontap-get...")
									return nil
								},
							},
							{
								Name:  "execute-ontap-delete",
								Usage: "execute-ontap-delete ontap",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing execute-ontap-delete...")
									return nil
								},
							},
							{
								Name:  "execute-ontap-patch",
								Usage: "execute-ontap-patch ontap",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing execute-ontap-patch...")
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
									client, err := netapp.NewClient(ctx)
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
									client, err := netapp.NewClient(ctx)
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
									client, err := netapp.NewClient(ctx)
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
									client, err := netapp.NewClient(ctx)
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
						Name:  "quota-rules",
						Usage: "Manage quota-rules resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list quota-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListQuotaRulesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListQuotaRules(ctx, req)
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
								Usage: "describe quota-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "quota_rule", Usage: "The quota_rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/quotaRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("quota_rule"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetQuotaRuleRequest{Name: name}
									resp, err := client.GetQuotaRule(ctx, req)
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
								Usage: "create quota-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "quota-rule-id", Usage: "The quota rule id.", Required: true},
									&cli.StringFlag{Name: "target", Usage: "The target.", Required: false},
									&cli.IntFlag{Name: "disk-limit-mib", Usage: "The disk limit mib.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateQuotaRuleRequest{Parent: parent}
									req.QuotaRuleId = cmd.String("quota-rule-id")
									req.QuotaRule = &netapppb.QuotaRule{
										Target:       cmd.String("target"),
										DiskLimitMib: int32(cmd.Int("disk-limit-mib")),
										Description:  cmd.String("description"),
									}
									op, err := client.CreateQuotaRule(ctx, req)
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
								Usage: "update quota-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "quota_rule", Usage: "The quota_rule.", Required: true},
									&cli.StringFlag{Name: "target", Usage: "The target.", Required: false},
									&cli.IntFlag{Name: "disk-limit-mib", Usage: "The disk limit mib.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/quotaRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("quota_rule"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateQuotaRuleRequest{}
									req.QuotaRule = &netapppb.QuotaRule{
										Name:         name,
										Target:       cmd.String("target"),
										DiskLimitMib: int32(cmd.Int("disk-limit-mib")),
										Description:  cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("target") {
										paths = append(paths, "target")
									}
									if cmd.IsSet("disk-limit-mib") {
										paths = append(paths, "disk_limit_mib")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateQuotaRule(ctx, req)
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
								Usage: "delete quota-rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "quota_rule", Usage: "The quota_rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/quotaRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("quota_rule"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteQuotaRuleRequest{Name: name}
									op, err := client.DeleteQuotaRule(ctx, req)
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
						Name:  "replications",
						Usage: "Manage replications resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListReplicationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListReplications(ctx, req)
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
								Usage: "describe replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetReplicationRequest{Name: name}
									resp, err := client.GetReplication(ctx, req)
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
								Usage: "create replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication-id", Usage: "The replication id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "cluster-location", Usage: "The cluster location.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateReplicationRequest{Parent: parent}
									req.ReplicationId = cmd.String("replication-id")
									req.Replication = &netapppb.Replication{
										Description:     cmd.String("description"),
										ClusterLocation: cmd.String("cluster-location"),
									}
									op, err := client.CreateReplication(ctx, req)
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
								Usage: "delete replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteReplicationRequest{Name: name}
									op, err := client.DeleteReplication(ctx, req)
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
								Usage: "update replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "cluster-location", Usage: "The cluster location.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateReplicationRequest{}
									req.Replication = &netapppb.Replication{
										Name:            name,
										Description:     cmd.String("description"),
										ClusterLocation: cmd.String("cluster-location"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("cluster-location") {
										paths = append(paths, "cluster_location")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateReplication(ctx, req)
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
								Name:  "stop",
								Usage: "stop replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
									&cli.BoolFlag{Name: "force", Usage: "The force.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.StopReplicationRequest{Name: name}
									req.Force = cmd.Bool("force")
									op, err := client.StopReplication(ctx, req)
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
								Name:  "resume",
								Usage: "resume replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.ResumeReplicationRequest{Name: name}
									op, err := client.ResumeReplication(ctx, req)
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
								Name:  "reverse-direction",
								Usage: "reverse-direction replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.ReverseReplicationDirectionRequest{Name: name}
									op, err := client.ReverseReplicationDirection(ctx, req)
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
								Name:  "establish-peering",
								Usage: "establish-peering replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
									&cli.StringFlag{Name: "peer-cluster-name", Usage: "The peer cluster name.", Required: true},
									&cli.StringFlag{Name: "peer-svm-name", Usage: "The peer svm name.", Required: true},
									&cli.StringFlag{Name: "peer-volume-name", Usage: "The peer volume name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.EstablishPeeringRequest{Name: name}
									req.PeerClusterName = cmd.String("peer-cluster-name")
									req.PeerSvmName = cmd.String("peer-svm-name")
									req.PeerVolumeName = cmd.String("peer-volume-name")
									op, err := client.EstablishPeering(ctx, req)
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
								Name:  "sync",
								Usage: "sync replications",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "replication", Usage: "The replication.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/replications/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("replication"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.SyncReplicationRequest{Name: name}
									op, err := client.SyncReplication(ctx, req)
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
						Name:  "snapshots",
						Usage: "Manage snapshots resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListSnapshotsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSnapshots(ctx, req)
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
								Usage: "describe snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetSnapshotRequest{Name: name}
									resp, err := client.GetSnapshot(ctx, req)
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
								Usage: "create snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "snapshot-id", Usage: "The snapshot id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateSnapshotRequest{Parent: parent}
									req.SnapshotId = cmd.String("snapshot-id")
									req.Snapshot = &netapppb.Snapshot{
										Description: cmd.String("description"),
									}
									op, err := client.CreateSnapshot(ctx, req)
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
								Usage: "delete snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteSnapshotRequest{Name: name}
									op, err := client.DeleteSnapshot(ctx, req)
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
								Usage: "update snapshots",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "snapshot", Usage: "The snapshot.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s/snapshots/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"), cmd.String("snapshot"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateSnapshotRequest{}
									req.Snapshot = &netapppb.Snapshot{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateSnapshot(ctx, req)
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
						Name:  "storage-pools",
						Usage: "Manage storage-pools resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list storage-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListStoragePoolsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListStoragePools(ctx, req)
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
								Usage: "create storage-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "storage-pool-id", Usage: "The storage pool id.", Required: true},
									&cli.IntFlag{Name: "capacity-gib", Usage: "The capacity gib.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
									&cli.StringFlag{Name: "active-directory", Usage: "The active directory.", Required: false},
									&cli.StringFlag{Name: "kms-config", Usage: "The kms config.", Required: false},
									&cli.BoolFlag{Name: "ldap-enabled", Usage: "The ldap enabled.", Required: false},
									&cli.StringFlag{Name: "psa-range", Usage: "The psa range.", Required: false},
									&cli.BoolFlag{Name: "global-access-allowed", Usage: "The global access allowed.", Required: false},
									&cli.BoolFlag{Name: "allow-auto-tiering", Usage: "The allow auto tiering.", Required: false},
									&cli.StringFlag{Name: "replica-zone", Usage: "The replica zone.", Required: false},
									&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: false},
									&cli.BoolFlag{Name: "custom-performance-enabled", Usage: "The custom performance enabled.", Required: false},
									&cli.IntFlag{Name: "total-throughput-mibps", Usage: "The total throughput mibps.", Required: false},
									&cli.IntFlag{Name: "total-iops", Usage: "The total iops.", Required: false},
									&cli.IntFlag{Name: "hot-tier-size-gib", Usage: "The hot tier size gib.", Required: false},
									&cli.BoolFlag{Name: "enable-hot-tier-auto-resize", Usage: "The enable hot tier auto resize.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateStoragePoolRequest{Parent: parent}
									req.StoragePoolId = cmd.String("storage-pool-id")
									req.StoragePool = &netapppb.StoragePool{
										CapacityGib:              int64(cmd.Int("capacity-gib")),
										Description:              cmd.String("description"),
										Network:                  cmd.String("network"),
										ActiveDirectory:          cmd.String("active-directory"),
										KmsConfig:                cmd.String("kms-config"),
										LdapEnabled:              cmd.Bool("ldap-enabled"),
										PsaRange:                 cmd.String("psa-range"),
										GlobalAccessAllowed:      cmd.Bool("global-access-allowed"),
										AllowAutoTiering:         cmd.Bool("allow-auto-tiering"),
										ReplicaZone:              cmd.String("replica-zone"),
										Zone:                     cmd.String("zone"),
										CustomPerformanceEnabled: cmd.Bool("custom-performance-enabled"),
										TotalThroughputMibps:     int64(cmd.Int("total-throughput-mibps")),
										TotalIops:                int64(cmd.Int("total-iops")),
										HotTierSizeGib:           int64(cmd.Int("hot-tier-size-gib")),
										EnableHotTierAutoResize:  cmd.Bool("enable-hot-tier-auto-resize"),
									}
									op, err := client.CreateStoragePool(ctx, req)
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
								Name:  "describe",
								Usage: "describe storage-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "storage_pool", Usage: "The storage_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage_pool"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetStoragePoolRequest{Name: name}
									resp, err := client.GetStoragePool(ctx, req)
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
								Usage: "update storage-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "storage_pool", Usage: "The storage_pool.", Required: true},
									&cli.IntFlag{Name: "capacity-gib", Usage: "The capacity gib.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.StringFlag{Name: "active-directory", Usage: "The active directory.", Required: false},
									&cli.StringFlag{Name: "kms-config", Usage: "The kms config.", Required: false},
									&cli.BoolFlag{Name: "ldap-enabled", Usage: "The ldap enabled.", Required: false},
									&cli.StringFlag{Name: "psa-range", Usage: "The psa range.", Required: false},
									&cli.BoolFlag{Name: "global-access-allowed", Usage: "The global access allowed.", Required: false},
									&cli.BoolFlag{Name: "allow-auto-tiering", Usage: "The allow auto tiering.", Required: false},
									&cli.StringFlag{Name: "replica-zone", Usage: "The replica zone.", Required: false},
									&cli.StringFlag{Name: "zone", Usage: "The zone.", Required: false},
									&cli.BoolFlag{Name: "custom-performance-enabled", Usage: "The custom performance enabled.", Required: false},
									&cli.IntFlag{Name: "total-throughput-mibps", Usage: "The total throughput mibps.", Required: false},
									&cli.IntFlag{Name: "total-iops", Usage: "The total iops.", Required: false},
									&cli.IntFlag{Name: "hot-tier-size-gib", Usage: "The hot tier size gib.", Required: false},
									&cli.BoolFlag{Name: "enable-hot-tier-auto-resize", Usage: "The enable hot tier auto resize.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage_pool"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateStoragePoolRequest{}
									req.StoragePool = &netapppb.StoragePool{
										Name:                     name,
										CapacityGib:              int64(cmd.Int("capacity-gib")),
										Description:              cmd.String("description"),
										Network:                  cmd.String("network"),
										ActiveDirectory:          cmd.String("active-directory"),
										KmsConfig:                cmd.String("kms-config"),
										LdapEnabled:              cmd.Bool("ldap-enabled"),
										PsaRange:                 cmd.String("psa-range"),
										GlobalAccessAllowed:      cmd.Bool("global-access-allowed"),
										AllowAutoTiering:         cmd.Bool("allow-auto-tiering"),
										ReplicaZone:              cmd.String("replica-zone"),
										Zone:                     cmd.String("zone"),
										CustomPerformanceEnabled: cmd.Bool("custom-performance-enabled"),
										TotalThroughputMibps:     int64(cmd.Int("total-throughput-mibps")),
										TotalIops:                int64(cmd.Int("total-iops")),
										HotTierSizeGib:           int64(cmd.Int("hot-tier-size-gib")),
										EnableHotTierAutoResize:  cmd.Bool("enable-hot-tier-auto-resize"),
									}
									var paths []string
									if cmd.IsSet("capacity-gib") {
										paths = append(paths, "capacity_gib")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									if cmd.IsSet("active-directory") {
										paths = append(paths, "active_directory")
									}
									if cmd.IsSet("kms-config") {
										paths = append(paths, "kms_config")
									}
									if cmd.IsSet("ldap-enabled") {
										paths = append(paths, "ldap_enabled")
									}
									if cmd.IsSet("psa-range") {
										paths = append(paths, "psa_range")
									}
									if cmd.IsSet("global-access-allowed") {
										paths = append(paths, "global_access_allowed")
									}
									if cmd.IsSet("allow-auto-tiering") {
										paths = append(paths, "allow_auto_tiering")
									}
									if cmd.IsSet("replica-zone") {
										paths = append(paths, "replica_zone")
									}
									if cmd.IsSet("zone") {
										paths = append(paths, "zone")
									}
									if cmd.IsSet("custom-performance-enabled") {
										paths = append(paths, "custom_performance_enabled")
									}
									if cmd.IsSet("total-throughput-mibps") {
										paths = append(paths, "total_throughput_mibps")
									}
									if cmd.IsSet("total-iops") {
										paths = append(paths, "total_iops")
									}
									if cmd.IsSet("hot-tier-size-gib") {
										paths = append(paths, "hot_tier_size_gib")
									}
									if cmd.IsSet("enable-hot-tier-auto-resize") {
										paths = append(paths, "enable_hot_tier_auto_resize")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateStoragePool(ctx, req)
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
								Usage: "delete storage-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "storage_pool", Usage: "The storage_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage_pool"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteStoragePoolRequest{Name: name}
									op, err := client.DeleteStoragePool(ctx, req)
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
								Name:  "validate-directory-service",
								Usage: "validate-directory-service storage-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "storage_pool", Usage: "The storage_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage_pool"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.ValidateDirectoryServiceRequest{Name: name}
									op, err := client.ValidateDirectoryService(ctx, req)
									if err != nil {
										return err
									}
									if err := op.Wait(ctx); err != nil {
										return err
									}
									fmt.Printf("ValidateDirectoryService completed on %s\n", name)
									return nil
								},
							},
							{
								Name:  "switch",
								Usage: "switch storage-pools",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "storage_pool", Usage: "The storage_pool.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", cmd.String("project"), cmd.String("location"), cmd.String("storage_pool"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.SwitchActiveReplicaZoneRequest{Name: name}
									op, err := client.SwitchActiveReplicaZone(ctx, req)
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
						Name:  "volumes",
						Usage: "Manage volumes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &netapppb.ListVolumesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVolumes(ctx, req)
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
								Usage: "describe volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.GetVolumeRequest{Name: name}
									resp, err := client.GetVolume(ctx, req)
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
								Usage: "create volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume-id", Usage: "The volume id.", Required: true},
									&cli.StringFlag{Name: "share-name", Usage: "The share name.", Required: true},
									&cli.StringFlag{Name: "storage-pool", Usage: "The storage pool.", Required: true},
									&cli.IntFlag{Name: "capacity-gib", Usage: "The capacity gib.", Required: true},
									&cli.StringFlag{Name: "unix-permissions", Usage: "The unix permissions.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "snapshot-directory", Usage: "The snapshot directory.", Required: false},
									&cli.BoolFlag{Name: "kerberos-enabled", Usage: "The kerberos enabled.", Required: false},
									&cli.BoolFlag{Name: "large-capacity", Usage: "The large capacity.", Required: false},
									&cli.BoolFlag{Name: "multiple-endpoints", Usage: "The multiple endpoints.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.CreateVolumeRequest{Parent: parent}
									req.VolumeId = cmd.String("volume-id")
									req.Volume = &netapppb.Volume{
										ShareName:         cmd.String("share-name"),
										StoragePool:       cmd.String("storage-pool"),
										CapacityGib:       int64(cmd.Int("capacity-gib")),
										UnixPermissions:   cmd.String("unix-permissions"),
										Description:       cmd.String("description"),
										SnapshotDirectory: cmd.Bool("snapshot-directory"),
										KerberosEnabled:   cmd.Bool("kerberos-enabled"),
										LargeCapacity:     cmd.Bool("large-capacity"),
										MultipleEndpoints: cmd.Bool("multiple-endpoints"),
									}
									op, err := client.CreateVolume(ctx, req)
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
								Usage: "update volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "share-name", Usage: "The share name.", Required: false},
									&cli.StringFlag{Name: "storage-pool", Usage: "The storage pool.", Required: false},
									&cli.IntFlag{Name: "capacity-gib", Usage: "The capacity gib.", Required: false},
									&cli.StringFlag{Name: "unix-permissions", Usage: "The unix permissions.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "snapshot-directory", Usage: "The snapshot directory.", Required: false},
									&cli.BoolFlag{Name: "kerberos-enabled", Usage: "The kerberos enabled.", Required: false},
									&cli.BoolFlag{Name: "large-capacity", Usage: "The large capacity.", Required: false},
									&cli.BoolFlag{Name: "multiple-endpoints", Usage: "The multiple endpoints.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.UpdateVolumeRequest{}
									req.Volume = &netapppb.Volume{
										Name:              name,
										ShareName:         cmd.String("share-name"),
										StoragePool:       cmd.String("storage-pool"),
										CapacityGib:       int64(cmd.Int("capacity-gib")),
										UnixPermissions:   cmd.String("unix-permissions"),
										Description:       cmd.String("description"),
										SnapshotDirectory: cmd.Bool("snapshot-directory"),
										KerberosEnabled:   cmd.Bool("kerberos-enabled"),
										LargeCapacity:     cmd.Bool("large-capacity"),
										MultipleEndpoints: cmd.Bool("multiple-endpoints"),
									}
									var paths []string
									if cmd.IsSet("share-name") {
										paths = append(paths, "share_name")
									}
									if cmd.IsSet("storage-pool") {
										paths = append(paths, "storage_pool")
									}
									if cmd.IsSet("capacity-gib") {
										paths = append(paths, "capacity_gib")
									}
									if cmd.IsSet("unix-permissions") {
										paths = append(paths, "unix_permissions")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("snapshot-directory") {
										paths = append(paths, "snapshot_directory")
									}
									if cmd.IsSet("kerberos-enabled") {
										paths = append(paths, "kerberos_enabled")
									}
									if cmd.IsSet("large-capacity") {
										paths = append(paths, "large_capacity")
									}
									if cmd.IsSet("multiple-endpoints") {
										paths = append(paths, "multiple_endpoints")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateVolume(ctx, req)
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
								Usage: "delete volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.DeleteVolumeRequest{Name: name}
									op, err := client.DeleteVolume(ctx, req)
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
								Name:  "revert",
								Usage: "revert volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "snapshot-id", Usage: "The snapshot id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.RevertVolumeRequest{Name: name}
									req.SnapshotId = cmd.String("snapshot-id")
									op, err := client.RevertVolume(ctx, req)
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
								Name:  "establish-peering",
								Usage: "establish-peering volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "peer-cluster-name", Usage: "The peer cluster name.", Required: true},
									&cli.StringFlag{Name: "peer-svm-name", Usage: "The peer svm name.", Required: true},
									&cli.StringFlag{Name: "peer-volume-name", Usage: "The peer volume name.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.EstablishVolumePeeringRequest{Name: name}
									req.PeerClusterName = cmd.String("peer-cluster-name")
									req.PeerSvmName = cmd.String("peer-svm-name")
									req.PeerVolumeName = cmd.String("peer-volume-name")
									op, err := client.EstablishVolumePeering(ctx, req)
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
								Name:  "restore",
								Usage: "restore volumes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "volume", Usage: "The volume.", Required: true},
									&cli.StringFlag{Name: "backup", Usage: "The backup.", Required: true},
									&cli.StringFlag{Name: "restore-destination-path", Usage: "The restore destination path.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/volumes/%s", cmd.String("project"), cmd.String("location"), cmd.String("volume"))
									client, err := netapp.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &netapppb.RestoreBackupFilesRequest{Name: name}
									req.Backup = cmd.String("backup")
									req.RestoreDestinationPath = cmd.String("restore-destination-path")
									op, err := client.RestoreBackupFiles(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
