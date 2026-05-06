package dlp

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud dlp command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "dlp",
		Usage: "manage Sensitive Data Protection (DLP) resources",
		Commands: []*cli.Command{
			{
				Name:  "column-data-profiles",
				Usage: "Manage column-data-profiles resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list column-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe column-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "column_data_profile", Usage: "The column_data_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/columnDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("column_data_profile"))
							fmt.Printf("Executing describe on %s\n", name)
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
						Name:  "search",
						Usage: "search connections",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing search on %s\n", parent)
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
				},
			},
			{
				Name:  "content",
				Usage: "Manage content resources",
				Commands: []*cli.Command{
					{
						Name:  "inspect",
						Usage: "inspect content",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/dlpContent", cmd.String("project"))
							fmt.Printf("Executing inspect on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "deidentify",
						Usage: "deidentify content",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/dlpContent", cmd.String("project"))
							fmt.Printf("Executing deidentify on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "reidentify",
						Usage: "reidentify content",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/dlpContent", cmd.String("project"))
							fmt.Printf("Executing reidentify on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "deidentify-templates",
				Usage: "Manage deidentify-templates resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "deidentify_template", Usage: "The deidentify_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/deidentifyTemplates/%s", cmd.String("organization"), cmd.String("deidentify_template"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "deidentify_template", Usage: "The deidentify_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/deidentifyTemplates/%s", cmd.String("organization"), cmd.String("deidentify_template"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete deidentify-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "deidentify_template", Usage: "The deidentify_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/deidentifyTemplates/%s", cmd.String("organization"), cmd.String("deidentify_template"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "discovery-configs",
				Usage: "Manage discovery-configs resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create discovery-configs",
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
						Usage: "update discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "discovery_config", Usage: "The discovery_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/discoveryConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("discovery_config"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "discovery_config", Usage: "The discovery_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/discoveryConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("discovery_config"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list discovery-configs",
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
						Name:  "delete",
						Usage: "delete discovery-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "discovery_config", Usage: "The discovery_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/discoveryConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("discovery_config"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "dlp-jobs",
				Usage: "Manage dlp-jobs resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create dlp-jobs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list dlp-jobs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp_job", Usage: "The dlp_job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp_job"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp_job", Usage: "The dlp_job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp_job"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "cancel",
						Usage: "cancel dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp_job", Usage: "The dlp_job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp_job"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
					{
						Name:  "hybrid-inspect",
						Usage: "hybrid-inspect dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp_job", Usage: "The dlp_job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp_job"))
							fmt.Printf("Executing hybrid-inspect on %s\n", name)
							return nil
						},
					},
					{
						Name:  "finish",
						Usage: "finish dlp-jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dlp_job", Usage: "The dlp_job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/dlpJobs/%s", cmd.String("project"), cmd.String("dlp_job"))
							fmt.Printf("Executing finish on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "file-store-data-profiles",
				Usage: "Manage file-store-data-profiles resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list file-store-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe file-store-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "project_data_profile", Usage: "The project_data_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/projectDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("project_data_profile"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete file-store-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "file_store_data_profile", Usage: "The file_store_data_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/fileStoreDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("file_store_data_profile"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "image",
				Usage: "Manage image resources",
				Commands: []*cli.Command{
					{
						Name:  "redact",
						Usage: "redact image",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/dlpContent", cmd.String("project"))
							fmt.Printf("Executing redact on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "info-types",
				Usage: "Manage info-types resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list info-types",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
			{
				Name:  "inspect-templates",
				Usage: "Manage inspect-templates resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "inspect_template", Usage: "The inspect_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/inspectTemplates/%s", cmd.String("organization"), cmd.String("inspect_template"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "inspect_template", Usage: "The inspect_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/inspectTemplates/%s", cmd.String("organization"), cmd.String("inspect_template"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete inspect-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "inspect_template", Usage: "The inspect_template.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/inspectTemplates/%s", cmd.String("organization"), cmd.String("inspect_template"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "job-triggers",
				Usage: "Manage job-triggers resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create job-triggers",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_trigger", Usage: "The job_trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job_trigger"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "hybrid-inspect",
						Usage: "hybrid-inspect job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_trigger", Usage: "The job_trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job_trigger"))
							fmt.Printf("Executing hybrid-inspect on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_trigger", Usage: "The job_trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job_trigger"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list job-triggers",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_trigger", Usage: "The job_trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job_trigger"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "activate",
						Usage: "activate job-triggers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_trigger", Usage: "The job_trigger.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/jobTriggers/%s", cmd.String("project"), cmd.String("job_trigger"))
							fmt.Printf("Executing activate on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "project-data-profiles",
				Usage: "Manage project-data-profiles resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list project-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe project-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "project_data_profile", Usage: "The project_data_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/projectDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("project_data_profile"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "stored-info-types",
				Usage: "Manage stored-info-types resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "stored_info_type", Usage: "The stored_info_type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/storedInfoTypes/%s", cmd.String("organization"), cmd.String("stored_info_type"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "stored_info_type", Usage: "The stored_info_type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/storedInfoTypes/%s", cmd.String("organization"), cmd.String("stored_info_type"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s", cmd.String("organization"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete stored-info-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "stored_info_type", Usage: "The stored_info_type.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/storedInfoTypes/%s", cmd.String("organization"), cmd.String("stored_info_type"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "table-data-profiles",
				Usage: "Manage table-data-profiles resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list table-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe table-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "table_data_profile", Usage: "The table_data_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/tableDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("table_data_profile"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete table-data-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "table_data_profile", Usage: "The table_data_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/tableDataProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("table_data_profile"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
