package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "monitoring",
				Usage: "manage Cloud Monitoring API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "alert-policies",
						Usage: "Manage alert-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list alert-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe alert-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "alert_policy", Usage: "The alert_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/alertPolicies/%s", cmd.String("project"), cmd.String("alert_policy"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create alert-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete alert-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "alert_policy", Usage: "The alert_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/alertPolicies/%s", cmd.String("project"), cmd.String("alert_policy"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update alert-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "alert_policy", Usage: "The alert_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/alertPolicies/%s", cmd.String("project"), cmd.String("alert_policy"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "groups",
						Usage: "Manage groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/groups/%s", cmd.String("project"), cmd.String("group"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/groups/%s", cmd.String("project"), cmd.String("group"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "group", Usage: "The group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/groups/%s", cmd.String("project"), cmd.String("group"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "members",
						Usage: "Manage members resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list members",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
						},
					},
					{
						Name:  "metric-descriptors",
						Usage: "Manage metric-descriptors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list metric-descriptors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe metric-descriptors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "metric_descriptor", Usage: "The metric_descriptor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/metricDescriptors/%s", cmd.String("project"), cmd.String("metric_descriptor"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create metric-descriptors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete metric-descriptors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "metric_descriptor", Usage: "The metric_descriptor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/metricDescriptors/%s", cmd.String("project"), cmd.String("metric_descriptor"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "monitored-resource-descriptors",
						Usage: "Manage monitored-resource-descriptors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list monitored-resource-descriptors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe monitored-resource-descriptors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "monitored_resource_descriptor", Usage: "The monitored_resource_descriptor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/monitoredResourceDescriptors/%s", cmd.String("project"), cmd.String("monitored_resource_descriptor"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "notification-channel-descriptors",
						Usage: "Manage notification-channel-descriptors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list notification-channel-descriptors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe notification-channel-descriptors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "channel_descriptor", Usage: "The channel_descriptor.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/notificationChannelDescriptors/%s", cmd.String("project"), cmd.String("channel_descriptor"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "notification-channels",
						Usage: "Manage notification-channels resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list notification-channels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe notification-channels",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "notification_channel", Usage: "The notification_channel.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification_channel"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create notification-channels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update notification-channels",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "notification_channel", Usage: "The notification_channel.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification_channel"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete notification-channels",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "notification_channel", Usage: "The notification_channel.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification_channel"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
							{
								Name:  "send-verification-code",
								Usage: "send-verification-code notification-channels",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "notification_channel", Usage: "The notification_channel.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification_channel"))
									fmt.Printf("Executing send-verification-code on %s\n", name)
									return nil
								},
							},
							{
								Name:  "get-verification-code",
								Usage: "get-verification-code notification-channels",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "notification_channel", Usage: "The notification_channel.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification_channel"))
									fmt.Printf("Executing get-verification-code on %s\n", name)
									return nil
								},
							},
							{
								Name:  "verify",
								Usage: "verify notification-channels",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "notification_channel", Usage: "The notification_channel.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/notificationChannels/%s", cmd.String("project"), cmd.String("notification_channel"))
									fmt.Printf("Executing verify on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "service-level-objectives",
						Usage: "Manage service-level-objectives resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create service-level-objectives",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/services/%s", cmd.String("project"), cmd.String("service"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe service-level-objectives",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "service_level_objective", Usage: "The service_level_objective.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/services/%s/serviceLevelObjectives/%s", cmd.String("project"), cmd.String("service"), cmd.String("service_level_objective"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list service-level-objectives",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update service-level-objectives",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "service_level_objective", Usage: "The service_level_objective.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/services/%s/serviceLevelObjectives/%s", cmd.String("project"), cmd.String("service"), cmd.String("service_level_objective"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete service-level-objectives",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.StringFlag{Name: "service_level_objective", Usage: "The service_level_objective.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/services/%s/serviceLevelObjectives/%s", cmd.String("project"), cmd.String("service"), cmd.String("service_level_objective"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "services",
						Usage: "Manage services resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/services/%s", cmd.String("project"), cmd.String("service"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/services/%s", cmd.String("project"), cmd.String("service"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete services",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/services/%s", cmd.String("project"), cmd.String("service"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "snoozes",
						Usage: "Manage snoozes resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create snoozes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list snoozes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe snoozes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "snooze", Usage: "The snooze.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/snoozes/%s", cmd.String("project"), cmd.String("snooze"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update snoozes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "snooze", Usage: "The snooze.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/snoozes/%s", cmd.String("project"), cmd.String("snooze"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "time-series",
						Usage: "Manage time-series resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list time-series",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create time-series",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create time-series",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "query",
								Usage: "query time-series",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing query...")
									return nil
								},
							},
						},
					},
					{
						Name:  "uptime-check-configs",
						Usage: "Manage uptime-check-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list uptime-check-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing list on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe uptime-check-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "uptime_check_config", Usage: "The uptime_check_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", cmd.String("project"), cmd.String("uptime_check_config"))
									fmt.Printf("Executing describe on %s\n", name)
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create uptime-check-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									fmt.Printf("Executing create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update uptime-check-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "uptime_check_config", Usage: "The uptime_check_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", cmd.String("project"), cmd.String("uptime_check_config"))
									fmt.Printf("Executing update on %s\n", name)
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete uptime-check-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "uptime_check_config", Usage: "The uptime_check_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/uptimeCheckConfigs/%s", cmd.String("project"), cmd.String("uptime_check_config"))
									fmt.Printf("Executing delete on %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "uptime-check-ips",
						Usage: "Manage uptime-check-ips resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list uptime-check-ips",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
