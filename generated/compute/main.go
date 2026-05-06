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
				Name:  "compute",
				Usage: "manage Google Compute Engine API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "abandon-instances",
						Usage: "Manage abandon-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "abandon-instances",
								Usage: "abandon-instances abandon-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing abandon-instances...")
									return nil
								},
							},
							{
								Name:  "abandon-instances",
								Usage: "abandon-instances abandon-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing abandon-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "accelerator-types",
						Usage: "Manage accelerator-types resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list accelerator-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe accelerator-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list accelerator-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-access-config",
						Usage: "Manage add-access-config resources",
						Commands: []*cli.Command{
							{
								Name:  "add-access-config",
								Usage: "add-access-config add-access-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-access-config...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-association",
						Usage: "Manage add-association resources",
						Commands: []*cli.Command{
							{
								Name:  "add-association",
								Usage: "add-association add-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-association...")
									return nil
								},
							},
							{
								Name:  "add-association",
								Usage: "add-association add-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-association...")
									return nil
								},
							},
							{
								Name:  "add-association",
								Usage: "add-association add-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-association...")
									return nil
								},
							},
							{
								Name:  "add-association",
								Usage: "add-association add-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-association...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-health-check",
						Usage: "Manage add-health-check resources",
						Commands: []*cli.Command{
							{
								Name:  "add-health-check",
								Usage: "add-health-check add-health-check",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-health-check...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-instance",
						Usage: "Manage add-instance resources",
						Commands: []*cli.Command{
							{
								Name:  "add-instance",
								Usage: "add-instance add-instance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-instance...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-instances",
						Usage: "Manage add-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "add-instances",
								Usage: "add-instances add-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-network-interface",
						Usage: "Manage add-network-interface resources",
						Commands: []*cli.Command{
							{
								Name:  "add-network-interface",
								Usage: "add-network-interface add-network-interface",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-network-interface...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-nodes",
						Usage: "Manage add-nodes resources",
						Commands: []*cli.Command{
							{
								Name:  "add-nodes",
								Usage: "add-nodes add-nodes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-nodes...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-packet-mirroring-rule",
						Usage: "Manage add-packet-mirroring-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "add-packet-mirroring-rule",
								Usage: "add-packet-mirroring-rule add-packet-mirroring-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-packet-mirroring-rule...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-peering",
						Usage: "Manage add-peering resources",
						Commands: []*cli.Command{
							{
								Name:  "add-peering",
								Usage: "add-peering add-peering",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-peering...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-resource-policies",
						Usage: "Manage add-resource-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "add-resource-policies",
								Usage: "add-resource-policies add-resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-resource-policies...")
									return nil
								},
							},
							{
								Name:  "add-resource-policies",
								Usage: "add-resource-policies add-resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-resource-policies...")
									return nil
								},
							},
							{
								Name:  "add-resource-policies",
								Usage: "add-resource-policies add-resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-resource-policies...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-rule",
						Usage: "Manage add-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "add-rule",
								Usage: "add-rule add-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-rule...")
									return nil
								},
							},
							{
								Name:  "add-rule",
								Usage: "add-rule add-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-rule...")
									return nil
								},
							},
							{
								Name:  "add-rule",
								Usage: "add-rule add-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-rule...")
									return nil
								},
							},
							{
								Name:  "add-rule",
								Usage: "add-rule add-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-rule...")
									return nil
								},
							},
							{
								Name:  "add-rule",
								Usage: "add-rule add-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-rule...")
									return nil
								},
							},
							{
								Name:  "add-rule",
								Usage: "add-rule add-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-rule...")
									return nil
								},
							},
						},
					},
					{
						Name:  "add-signed-url-key",
						Usage: "Manage add-signed-url-key resources",
						Commands: []*cli.Command{
							{
								Name:  "add-signed-url-key",
								Usage: "add-signed-url-key add-signed-url-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-signed-url-key...")
									return nil
								},
							},
							{
								Name:  "add-signed-url-key",
								Usage: "add-signed-url-key add-signed-url-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-signed-url-key...")
									return nil
								},
							},
						},
					},
					{
						Name:  "addresses",
						Usage: "Manage addresses resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list addresses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "announce",
						Usage: "Manage announce resources",
						Commands: []*cli.Command{
							{
								Name:  "announce",
								Usage: "announce announce",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing announce...")
									return nil
								},
							},
							{
								Name:  "announce",
								Usage: "announce announce",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing announce...")
									return nil
								},
							},
						},
					},
					{
						Name:  "apply-updates-to-instances",
						Usage: "Manage apply-updates-to-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "apply-updates-to-instances",
								Usage: "apply-updates-to-instances apply-updates-to-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing apply-updates-to-instances...")
									return nil
								},
							},
							{
								Name:  "apply-updates-to-instances",
								Usage: "apply-updates-to-instances apply-updates-to-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing apply-updates-to-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "attach-disk",
						Usage: "Manage attach-disk resources",
						Commands: []*cli.Command{
							{
								Name:  "attach-disk",
								Usage: "attach-disk attach-disk",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing attach-disk...")
									return nil
								},
							},
						},
					},
					{
						Name:  "attach-network-endpoints",
						Usage: "Manage attach-network-endpoints resources",
						Commands: []*cli.Command{
							{
								Name:  "attach-network-endpoints",
								Usage: "attach-network-endpoints attach-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing attach-network-endpoints...")
									return nil
								},
							},
							{
								Name:  "attach-network-endpoints",
								Usage: "attach-network-endpoints attach-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing attach-network-endpoints...")
									return nil
								},
							},
							{
								Name:  "attach-network-endpoints",
								Usage: "attach-network-endpoints attach-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing attach-network-endpoints...")
									return nil
								},
							},
						},
					},
					{
						Name:  "autoscalers",
						Usage: "Manage autoscalers resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update autoscalers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "backend-buckets",
						Usage: "Manage backend-buckets resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch backend-buckets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "backend-services",
						Usage: "Manage backend-services resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update backend-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "bulk-insert",
						Usage: "Manage bulk-insert resources",
						Commands: []*cli.Command{
							{
								Name:  "bulk-insert",
								Usage: "bulk-insert bulk-insert",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing bulk-insert...")
									return nil
								},
							},
							{
								Name:  "bulk-insert",
								Usage: "bulk-insert bulk-insert",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing bulk-insert...")
									return nil
								},
							},
							{
								Name:  "bulk-insert",
								Usage: "bulk-insert bulk-insert",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing bulk-insert...")
									return nil
								},
							},
							{
								Name:  "bulk-insert",
								Usage: "bulk-insert bulk-insert",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing bulk-insert...")
									return nil
								},
							},
						},
					},
					{
						Name:  "bulk-set-labels",
						Usage: "Manage bulk-set-labels resources",
						Commands: []*cli.Command{
							{
								Name:  "bulk-set-labels",
								Usage: "bulk-set-labels bulk-set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing bulk-set-labels...")
									return nil
								},
							},
						},
					},
					{
						Name:  "calendar-mode",
						Usage: "Manage calendar-mode resources",
						Commands: []*cli.Command{
							{
								Name:  "calendar-mode",
								Usage: "calendar-mode calendar-mode",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing calendar-mode...")
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
							{
								Name:  "cancel",
								Usage: "cancel cancel",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing cancel...")
									return nil
								},
							},
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
						Name:  "clone-rules",
						Usage: "Manage clone-rules resources",
						Commands: []*cli.Command{
							{
								Name:  "clone-rules",
								Usage: "clone-rules clone-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing clone-rules...")
									return nil
								},
							},
							{
								Name:  "clone-rules",
								Usage: "clone-rules clone-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing clone-rules...")
									return nil
								},
							},
							{
								Name:  "clone-rules",
								Usage: "clone-rules clone-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing clone-rules...")
									return nil
								},
							},
						},
					},
					{
						Name:  "commitments",
						Usage: "Manage commitments resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list commitments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe commitments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert commitments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list commitments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update commitments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "composite-health-checks",
						Usage: "Manage composite-health-checks resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list composite-health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete composite-health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe composite-health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert composite-health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list composite-health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch composite-health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "copy-rules",
						Usage: "Manage copy-rules resources",
						Commands: []*cli.Command{
							{
								Name:  "copy-rules",
								Usage: "copy-rules copy-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing copy-rules...")
									return nil
								},
							},
						},
					},
					{
						Name:  "create-instances",
						Usage: "Manage create-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create create-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create create-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
						},
					},
					{
						Name:  "create-members",
						Usage: "Manage create-members resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create create-members",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
						},
					},
					{
						Name:  "create-snapshot",
						Usage: "Manage create-snapshot resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create create-snapshot",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create create-snapshot",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
						},
					},
					{
						Name:  "cross-site-networks",
						Usage: "Manage cross-site-networks resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete cross-site-networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe cross-site-networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert cross-site-networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list cross-site-networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch cross-site-networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete-access-config",
						Usage: "Manage delete-access-config resources",
						Commands: []*cli.Command{
							{
								Name:  "delete-access-config",
								Usage: "delete-access-config delete-access-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-access-config...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete-instances",
						Usage: "Manage delete-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "delete-instances",
								Usage: "delete-instances delete-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-instances...")
									return nil
								},
							},
							{
								Name:  "delete-instances",
								Usage: "delete-instances delete-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete-network-interface",
						Usage: "Manage delete-network-interface resources",
						Commands: []*cli.Command{
							{
								Name:  "delete-network-interface",
								Usage: "delete-network-interface delete-network-interface",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-network-interface...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete-nodes",
						Usage: "Manage delete-nodes resources",
						Commands: []*cli.Command{
							{
								Name:  "delete-nodes",
								Usage: "delete-nodes delete-nodes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-nodes...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete-per-instance-configs",
						Usage: "Manage delete-per-instance-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "delete-per-instance-configs",
								Usage: "delete-per-instance-configs delete-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-per-instance-configs...")
									return nil
								},
							},
							{
								Name:  "delete-per-instance-configs",
								Usage: "delete-per-instance-configs delete-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-per-instance-configs...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete-route-policy",
						Usage: "Manage delete-route-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "delete-route-policy",
								Usage: "delete-route-policy delete-route-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-route-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "delete-signed-url-key",
						Usage: "Manage delete-signed-url-key resources",
						Commands: []*cli.Command{
							{
								Name:  "delete-signed-url-key",
								Usage: "delete-signed-url-key delete-signed-url-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-signed-url-key...")
									return nil
								},
							},
							{
								Name:  "delete-signed-url-key",
								Usage: "delete-signed-url-key delete-signed-url-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete-signed-url-key...")
									return nil
								},
							},
						},
					},
					{
						Name:  "deprecate",
						Usage: "Manage deprecate resources",
						Commands: []*cli.Command{
							{
								Name:  "deprecate",
								Usage: "deprecate deprecate",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing deprecate...")
									return nil
								},
							},
						},
					},
					{
						Name:  "detach-disk",
						Usage: "Manage detach-disk resources",
						Commands: []*cli.Command{
							{
								Name:  "detach-disk",
								Usage: "detach-disk detach-disk",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing detach-disk...")
									return nil
								},
							},
						},
					},
					{
						Name:  "detach-network-endpoints",
						Usage: "Manage detach-network-endpoints resources",
						Commands: []*cli.Command{
							{
								Name:  "detach-network-endpoints",
								Usage: "detach-network-endpoints detach-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing detach-network-endpoints...")
									return nil
								},
							},
							{
								Name:  "detach-network-endpoints",
								Usage: "detach-network-endpoints detach-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing detach-network-endpoints...")
									return nil
								},
							},
							{
								Name:  "detach-network-endpoints",
								Usage: "detach-network-endpoints detach-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing detach-network-endpoints...")
									return nil
								},
							},
						},
					},
					{
						Name:  "disable-xpn-host",
						Usage: "Manage disable-xpn-host resources",
						Commands: []*cli.Command{
							{
								Name:  "disable-xpn-host",
								Usage: "disable-xpn-host disable-xpn-host",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing disable-xpn-host...")
									return nil
								},
							},
						},
					},
					{
						Name:  "disable-xpn-resource",
						Usage: "Manage disable-xpn-resource resources",
						Commands: []*cli.Command{
							{
								Name:  "disable-xpn-resource",
								Usage: "disable-xpn-resource disable-xpn-resource",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing disable-xpn-resource...")
									return nil
								},
							},
						},
					},
					{
						Name:  "disk-types",
						Usage: "Manage disk-types resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list disk-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe disk-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list disk-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe disk-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list disk-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "disks",
						Usage: "Manage disks resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "enable-xpn-host",
						Usage: "Manage enable-xpn-host resources",
						Commands: []*cli.Command{
							{
								Name:  "enable-xpn-host",
								Usage: "enable-xpn-host enable-xpn-host",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing enable-xpn-host...")
									return nil
								},
							},
						},
					},
					{
						Name:  "enable-xpn-resource",
						Usage: "Manage enable-xpn-resource resources",
						Commands: []*cli.Command{
							{
								Name:  "enable-xpn-resource",
								Usage: "enable-xpn-resource enable-xpn-resource",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing enable-xpn-resource...")
									return nil
								},
							},
						},
					},
					{
						Name:  "expand-ip-cidr-range",
						Usage: "Manage expand-ip-cidr-range resources",
						Commands: []*cli.Command{
							{
								Name:  "expand-ip-cidr-range",
								Usage: "expand-ip-cidr-range expand-ip-cidr-range",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing expand-ip-cidr-range...")
									return nil
								},
							},
						},
					},
					{
						Name:  "external-vpn-gateways",
						Usage: "Manage external-vpn-gateways resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete external-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe external-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert external-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list external-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "family",
						Usage: "Manage family resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe family",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "firewall-policies",
						Usage: "Manage firewall-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch firewall-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "firewalls",
						Usage: "Manage firewalls resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "forwarding-rules",
						Usage: "Manage forwarding-rules resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch forwarding-rules",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "future-reservations",
						Usage: "Manage future-reservations resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list future-reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete future-reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe future-reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert future-reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list future-reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update future-reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-association",
						Usage: "Manage get-association resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-diagnostics",
						Usage: "Manage get-diagnostics resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-diagnostics",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-effective-firewalls",
						Usage: "Manage get-effective-firewalls resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-effective-firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-effective-firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-effective-firewalls",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-effective-security-policies",
						Usage: "Manage get-effective-security-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-effective-security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-guest-attributes",
						Usage: "Manage get-guest-attributes resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-guest-attributes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-health",
						Usage: "Manage get-health resources",
						Commands: []*cli.Command{
							{
								Name:  "get-health",
								Usage: "get-health get-health",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-health...")
									return nil
								},
							},
							{
								Name:  "get-health",
								Usage: "get-health get-health",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-health...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-health",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-health",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "get-health",
								Usage: "get-health get-health",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-health...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "Manage get-iam-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-macsec-config",
						Usage: "Manage get-macsec-config resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-macsec-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-nat-ip-info",
						Usage: "Manage get-nat-ip-info resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-nat-ip-info",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-nat-mapping-info",
						Usage: "Manage get-nat-mapping-info resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-nat-mapping-info",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-operational-status",
						Usage: "Manage get-operational-status resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-operational-status",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-operational-status",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-packet-mirroring-rule",
						Usage: "Manage get-packet-mirroring-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-packet-mirroring-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-route-policy",
						Usage: "Manage get-route-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-route-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-router-status",
						Usage: "Manage get-router-status resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-router-status",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-rule",
						Usage: "Manage get-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe get-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-shielded-instance-identity",
						Usage: "Manage get-shielded-instance-identity resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-shielded-instance-identity",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-status",
						Usage: "Manage get-status resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-status",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-version",
						Usage: "Manage get-version resources",
						Commands: []*cli.Command{
							{
								Name:  "get-version",
								Usage: "get-version get-version",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-version...")
									return nil
								},
							},
							{
								Name:  "get-version",
								Usage: "get-version get-version",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-version...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-xpn-host",
						Usage: "Manage get-xpn-host resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-xpn-host",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "get-xpn-resources",
						Usage: "Manage get-xpn-resources resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe get-xpn-resources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "health-aggregation-policies",
						Usage: "Manage health-aggregation-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list health-aggregation-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete health-aggregation-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe health-aggregation-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert health-aggregation-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list health-aggregation-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch health-aggregation-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "health-check-services",
						Usage: "Manage health-check-services resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list health-check-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete health-check-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe health-check-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert health-check-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list health-check-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch health-check-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "health-checks",
						Usage: "Manage health-checks resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update health-checks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "health-sources",
						Usage: "Manage health-sources resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list health-sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete health-sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe health-sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert health-sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list health-sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch health-sources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "image-family-views",
						Usage: "Manage image-family-views resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe image-family-views",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "images",
						Usage: "Manage images resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "instance-group-managers",
						Usage: "Manage instance-group-managers resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch instance-group-managers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "instance-groups",
						Usage: "Manage instance-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list instance-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instance-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instance-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instance-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instance-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instance-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instance-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "instance-settings",
						Usage: "Manage instance-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe instance-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch instance-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "instance-templates",
						Usage: "Manage instance-templates resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instance-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
								Name:  "aggregated-list",
								Usage: "aggregated-list instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
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
						Name:  "instant-snapshot-groups",
						Usage: "Manage instant-snapshot-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instant-snapshot-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "instant-snapshots",
						Usage: "Manage instant-snapshots resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list instant-snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "interconnect-attachment-groups",
						Usage: "Manage interconnect-attachment-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete interconnect-attachment-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe interconnect-attachment-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert interconnect-attachment-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list interconnect-attachment-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch interconnect-attachment-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "interconnect-attachments",
						Usage: "Manage interconnect-attachments resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list interconnect-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete interconnect-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe interconnect-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert interconnect-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list interconnect-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch interconnect-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "interconnect-groups",
						Usage: "Manage interconnect-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete interconnect-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe interconnect-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert interconnect-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list interconnect-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch interconnect-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "interconnect-locations",
						Usage: "Manage interconnect-locations resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe interconnect-locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list interconnect-locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "interconnect-remote-locations",
						Usage: "Manage interconnect-remote-locations resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe interconnect-remote-locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list interconnect-remote-locations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "interconnects",
						Usage: "Manage interconnects resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete interconnects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe interconnects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert interconnects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list interconnects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch interconnects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "invalidate-cache",
						Usage: "Manage invalidate-cache resources",
						Commands: []*cli.Command{
							{
								Name:  "invalidate-cache",
								Usage: "invalidate-cache invalidate-cache",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing invalidate-cache...")
									return nil
								},
							},
						},
					},
					{
						Name:  "license-codes",
						Usage: "Manage license-codes resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe license-codes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "licenses",
						Usage: "Manage licenses resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete licenses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe licenses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert licenses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list licenses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update licenses",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-associations",
						Usage: "Manage list-associations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-associations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-associations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-available-features",
						Usage: "Manage list-available-features resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-available-features",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-available-features",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-bgp-routes",
						Usage: "Manage list-bgp-routes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-bgp-routes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-disks",
						Usage: "Manage list-disks resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-disks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-errors",
						Usage: "Manage list-errors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-errors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-errors",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-instances",
						Usage: "Manage list-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "list-instances",
								Usage: "list-instances list-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-instances...")
									return nil
								},
							},
							{
								Name:  "list-instances",
								Usage: "list-instances list-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-managed-instances",
						Usage: "Manage list-managed-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "list-managed-instances",
								Usage: "list-managed-instances list-managed-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-managed-instances...")
									return nil
								},
							},
							{
								Name:  "list-managed-instances",
								Usage: "list-managed-instances list-managed-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-managed-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-network-endpoints",
						Usage: "Manage list-network-endpoints resources",
						Commands: []*cli.Command{
							{
								Name:  "list-network-endpoints",
								Usage: "list-network-endpoints list-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-network-endpoints...")
									return nil
								},
							},
							{
								Name:  "list-network-endpoints",
								Usage: "list-network-endpoints list-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-network-endpoints...")
									return nil
								},
							},
							{
								Name:  "list-network-endpoints",
								Usage: "list-network-endpoints list-network-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-network-endpoints...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-nodes",
						Usage: "Manage list-nodes resources",
						Commands: []*cli.Command{
							{
								Name:  "list-nodes",
								Usage: "list-nodes list-nodes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-nodes...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-peering-routes",
						Usage: "Manage list-peering-routes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-peering-routes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-per-instance-configs",
						Usage: "Manage list-per-instance-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "list-per-instance-configs",
								Usage: "list-per-instance-configs list-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-per-instance-configs...")
									return nil
								},
							},
							{
								Name:  "list-per-instance-configs",
								Usage: "list-per-instance-configs list-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-per-instance-configs...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-preconfigured-expression-sets",
						Usage: "Manage list-preconfigured-expression-sets resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-preconfigured-expression-sets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-preconfigured-expression-sets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-route-policies",
						Usage: "Manage list-route-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-route-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-usable",
						Usage: "Manage list-usable resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list list-usable",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-usable",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-usable",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-usable",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list list-usable",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "list-xpn-hosts",
						Usage: "Manage list-xpn-hosts resources",
						Commands: []*cli.Command{
							{
								Name:  "list-xpn-hosts",
								Usage: "list-xpn-hosts list-xpn-hosts",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list-xpn-hosts...")
									return nil
								},
							},
						},
					},
					{
						Name:  "machine-images",
						Usage: "Manage machine-images resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete machine-images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe machine-images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert machine-images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list machine-images",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "machine-types",
						Usage: "Manage machine-types resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list machine-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe machine-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list machine-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "move",
						Usage: "Manage move resources",
						Commands: []*cli.Command{
							{
								Name:  "move",
								Usage: "move move",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move...")
									return nil
								},
							},
							{
								Name:  "move",
								Usage: "move move",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move...")
									return nil
								},
							},
							{
								Name:  "move",
								Usage: "move move",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move...")
									return nil
								},
							},
							{
								Name:  "move",
								Usage: "move move",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move...")
									return nil
								},
							},
						},
					},
					{
						Name:  "move-disk",
						Usage: "Manage move-disk resources",
						Commands: []*cli.Command{
							{
								Name:  "move-disk",
								Usage: "move-disk move-disk",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move-disk...")
									return nil
								},
							},
						},
					},
					{
						Name:  "move-instance",
						Usage: "Manage move-instance resources",
						Commands: []*cli.Command{
							{
								Name:  "move-instance",
								Usage: "move-instance move-instance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing move-instance...")
									return nil
								},
							},
						},
					},
					{
						Name:  "network-attachments",
						Usage: "Manage network-attachments resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list network-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete network-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe network-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert network-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list network-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch network-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "network-edge-security-services",
						Usage: "Manage network-edge-security-services resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list network-edge-security-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete network-edge-security-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe network-edge-security-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert network-edge-security-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch network-edge-security-services",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "network-endpoint-groups",
						Usage: "Manage network-endpoint-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list network-endpoint-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "network-profiles",
						Usage: "Manage network-profiles resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe network-profiles",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list network-profiles",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "networks",
						Usage: "Manage networks resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch networks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "node-groups",
						Usage: "Manage node-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list node-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete node-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe node-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert node-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list node-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch node-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "node-templates",
						Usage: "Manage node-templates resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list node-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete node-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe node-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert node-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list node-templates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "node-types",
						Usage: "Manage node-types resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list node-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe node-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list node-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "notification-endpoints",
						Usage: "Manage notification-endpoints resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list notification-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete notification-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe notification-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert notification-endpoints",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list notification-endpoints",
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
								Name:  "aggregated-list",
								Usage: "aggregated-list operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
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
							{
								Name:  "delete",
								Usage: "delete operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
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
							{
								Name:  "delete",
								Usage: "delete operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
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
							{
								Name:  "delete",
								Usage: "delete operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
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
						Name:  "packet-mirrorings",
						Usage: "Manage packet-mirrorings resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list packet-mirrorings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete packet-mirrorings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe packet-mirrorings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert packet-mirrorings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list packet-mirrorings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch packet-mirrorings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "patch-packet-mirroring-rule",
						Usage: "Manage patch-packet-mirroring-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "patch-packet-mirroring-rule",
								Usage: "patch-packet-mirroring-rule patch-packet-mirroring-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-packet-mirroring-rule...")
									return nil
								},
							},
						},
					},
					{
						Name:  "patch-per-instance-configs",
						Usage: "Manage patch-per-instance-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "patch-per-instance-configs",
								Usage: "patch-per-instance-configs patch-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-per-instance-configs...")
									return nil
								},
							},
							{
								Name:  "patch-per-instance-configs",
								Usage: "patch-per-instance-configs patch-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-per-instance-configs...")
									return nil
								},
							},
						},
					},
					{
						Name:  "patch-route-policy",
						Usage: "Manage patch-route-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "patch-route-policy",
								Usage: "patch-route-policy patch-route-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-route-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "patch-rule",
						Usage: "Manage patch-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "patch-rule",
								Usage: "patch-rule patch-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-rule...")
									return nil
								},
							},
							{
								Name:  "patch-rule",
								Usage: "patch-rule patch-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-rule...")
									return nil
								},
							},
							{
								Name:  "patch-rule",
								Usage: "patch-rule patch-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-rule...")
									return nil
								},
							},
							{
								Name:  "patch-rule",
								Usage: "patch-rule patch-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-rule...")
									return nil
								},
							},
							{
								Name:  "patch-rule",
								Usage: "patch-rule patch-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-rule...")
									return nil
								},
							},
							{
								Name:  "patch-rule",
								Usage: "patch-rule patch-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch-rule...")
									return nil
								},
							},
						},
					},
					{
						Name:  "perform-maintenance",
						Usage: "Manage perform-maintenance resources",
						Commands: []*cli.Command{
							{
								Name:  "perform-maintenance",
								Usage: "perform-maintenance perform-maintenance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing perform-maintenance...")
									return nil
								},
							},
							{
								Name:  "perform-maintenance",
								Usage: "perform-maintenance perform-maintenance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing perform-maintenance...")
									return nil
								},
							},
							{
								Name:  "perform-maintenance",
								Usage: "perform-maintenance perform-maintenance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing perform-maintenance...")
									return nil
								},
							},
							{
								Name:  "perform-maintenance",
								Usage: "perform-maintenance perform-maintenance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing perform-maintenance...")
									return nil
								},
							},
							{
								Name:  "perform-maintenance",
								Usage: "perform-maintenance perform-maintenance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing perform-maintenance...")
									return nil
								},
							},
						},
					},
					{
						Name:  "preview",
						Usage: "Manage preview resources",
						Commands: []*cli.Command{
							{
								Name:  "preview",
								Usage: "preview preview",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing preview...")
									return nil
								},
							},
						},
					},
					{
						Name:  "preview-features",
						Usage: "Manage preview-features resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe preview-features",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list preview-features",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update preview-features",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
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
								Name:  "describe",
								Usage: "describe projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "public-advertised-prefixes",
						Usage: "Manage public-advertised-prefixes resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete public-advertised-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe public-advertised-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert public-advertised-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list public-advertised-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch public-advertised-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "public-delegated-prefixes",
						Usage: "Manage public-delegated-prefixes resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch public-delegated-prefixes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "recreate-instances",
						Usage: "Manage recreate-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "recreate-instances",
								Usage: "recreate-instances recreate-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing recreate-instances...")
									return nil
								},
							},
							{
								Name:  "recreate-instances",
								Usage: "recreate-instances recreate-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing recreate-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "referrers",
						Usage: "Manage referrers resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list referrers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "regions",
						Usage: "Manage regions resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe regions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list regions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-association",
						Usage: "Manage remove-association resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-association",
								Usage: "remove-association remove-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-association...")
									return nil
								},
							},
							{
								Name:  "remove-association",
								Usage: "remove-association remove-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-association...")
									return nil
								},
							},
							{
								Name:  "remove-association",
								Usage: "remove-association remove-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-association...")
									return nil
								},
							},
							{
								Name:  "remove-association",
								Usage: "remove-association remove-association",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-association...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-health-check",
						Usage: "Manage remove-health-check resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-health-check",
								Usage: "remove-health-check remove-health-check",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-health-check...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-instance",
						Usage: "Manage remove-instance resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-instance",
								Usage: "remove-instance remove-instance",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-instance...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-instances",
						Usage: "Manage remove-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-instances",
								Usage: "remove-instances remove-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-packet-mirroring-rule",
						Usage: "Manage remove-packet-mirroring-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-packet-mirroring-rule",
								Usage: "remove-packet-mirroring-rule remove-packet-mirroring-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-packet-mirroring-rule...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-peering",
						Usage: "Manage remove-peering resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-peering",
								Usage: "remove-peering remove-peering",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-peering...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-resource-policies",
						Usage: "Manage remove-resource-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-resource-policies",
								Usage: "remove-resource-policies remove-resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-resource-policies...")
									return nil
								},
							},
							{
								Name:  "remove-resource-policies",
								Usage: "remove-resource-policies remove-resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-resource-policies...")
									return nil
								},
							},
							{
								Name:  "remove-resource-policies",
								Usage: "remove-resource-policies remove-resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-resource-policies...")
									return nil
								},
							},
						},
					},
					{
						Name:  "remove-rule",
						Usage: "Manage remove-rule resources",
						Commands: []*cli.Command{
							{
								Name:  "remove-rule",
								Usage: "remove-rule remove-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-rule...")
									return nil
								},
							},
							{
								Name:  "remove-rule",
								Usage: "remove-rule remove-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-rule...")
									return nil
								},
							},
							{
								Name:  "remove-rule",
								Usage: "remove-rule remove-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-rule...")
									return nil
								},
							},
							{
								Name:  "remove-rule",
								Usage: "remove-rule remove-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-rule...")
									return nil
								},
							},
							{
								Name:  "remove-rule",
								Usage: "remove-rule remove-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-rule...")
									return nil
								},
							},
							{
								Name:  "remove-rule",
								Usage: "remove-rule remove-rule",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-rule...")
									return nil
								},
							},
						},
					},
					{
						Name:  "report-faulty",
						Usage: "Manage report-faulty resources",
						Commands: []*cli.Command{
							{
								Name:  "report-faulty",
								Usage: "report-faulty report-faulty",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing report-faulty...")
									return nil
								},
							},
						},
					},
					{
						Name:  "report-host-as-faulty",
						Usage: "Manage report-host-as-faulty resources",
						Commands: []*cli.Command{
							{
								Name:  "report-host-as-faulty",
								Usage: "report-host-as-faulty report-host-as-faulty",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing report-host-as-faulty...")
									return nil
								},
							},
						},
					},
					{
						Name:  "request-remove-peering",
						Usage: "Manage request-remove-peering resources",
						Commands: []*cli.Command{
							{
								Name:  "request-remove-peering",
								Usage: "request-remove-peering request-remove-peering",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing request-remove-peering...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reservation-blocks",
						Usage: "Manage reservation-blocks resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe reservation-blocks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list reservation-blocks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reservation-slots",
						Usage: "Manage reservation-slots resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe reservation-slots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list reservation-slots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update reservation-slots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reservation-sub-blocks",
						Usage: "Manage reservation-sub-blocks resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe reservation-sub-blocks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list reservation-sub-blocks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reservations",
						Usage: "Manage reservations resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update reservations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "reset",
						Usage: "Manage reset resources",
						Commands: []*cli.Command{
							{
								Name:  "reset",
								Usage: "reset reset",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing reset...")
									return nil
								},
							},
						},
					},
					{
						Name:  "resize",
						Usage: "Manage resize resources",
						Commands: []*cli.Command{
							{
								Name:  "resize",
								Usage: "resize resize",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resize...")
									return nil
								},
							},
							{
								Name:  "resize",
								Usage: "resize resize",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resize...")
									return nil
								},
							},
							{
								Name:  "resize",
								Usage: "resize resize",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resize...")
									return nil
								},
							},
							{
								Name:  "resize",
								Usage: "resize resize",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resize...")
									return nil
								},
							},
							{
								Name:  "resize",
								Usage: "resize resize",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resize...")
									return nil
								},
							},
						},
					},
					{
						Name:  "resize-requests",
						Usage: "Manage resize-requests resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list resize-requests",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "resource-policies",
						Usage: "Manage resource-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch resource-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "resume",
						Usage: "Manage resume resources",
						Commands: []*cli.Command{
							{
								Name:  "resume",
								Usage: "resume resume",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resume...")
									return nil
								},
							},
						},
					},
					{
						Name:  "resume-instances",
						Usage: "Manage resume-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "resume-instances",
								Usage: "resume-instances resume-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resume-instances...")
									return nil
								},
							},
							{
								Name:  "resume-instances",
								Usage: "resume-instances resume-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing resume-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "routers",
						Usage: "Manage routers resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list routers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete routers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe routers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert routers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list routers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch routers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update routers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "routes",
						Usage: "Manage routes resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete routes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe routes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert routes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list routes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "screenshot",
						Usage: "Manage screenshot resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe screenshot",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "security-policies",
						Usage: "Manage security-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "send-diagnostic-interrupt",
						Usage: "Manage send-diagnostic-interrupt resources",
						Commands: []*cli.Command{
							{
								Name:  "send-diagnostic-interrupt",
								Usage: "send-diagnostic-interrupt send-diagnostic-interrupt",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing send-diagnostic-interrupt...")
									return nil
								},
							},
						},
					},
					{
						Name:  "serial-port",
						Usage: "Manage serial-port resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe serial-port",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "service-attachments",
						Usage: "Manage service-attachments resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list service-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete service-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe service-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert service-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list service-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch service-attachments",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-backend-service",
						Usage: "Manage set-backend-service resources",
						Commands: []*cli.Command{
							{
								Name:  "set-backend-service",
								Usage: "set-backend-service set-backend-service",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-backend-service...")
									return nil
								},
							},
							{
								Name:  "set-backend-service",
								Usage: "set-backend-service set-backend-service",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-backend-service...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-backup",
						Usage: "Manage set-backup resources",
						Commands: []*cli.Command{
							{
								Name:  "set-backup",
								Usage: "set-backup set-backup",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-backup...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-certificate-map",
						Usage: "Manage set-certificate-map resources",
						Commands: []*cli.Command{
							{
								Name:  "set-certificate-map",
								Usage: "set-certificate-map set-certificate-map",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-certificate-map...")
									return nil
								},
							},
							{
								Name:  "set-certificate-map",
								Usage: "set-certificate-map set-certificate-map",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-certificate-map...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-cloud-armor-tier",
						Usage: "Manage set-cloud-armor-tier resources",
						Commands: []*cli.Command{
							{
								Name:  "set-cloud-armor-tier",
								Usage: "set-cloud-armor-tier set-cloud-armor-tier",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-cloud-armor-tier...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-common-instance-metadata",
						Usage: "Manage set-common-instance-metadata resources",
						Commands: []*cli.Command{
							{
								Name:  "set-common-instance-metadata",
								Usage: "set-common-instance-metadata set-common-instance-metadata",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-common-instance-metadata...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-default-network-tier",
						Usage: "Manage set-default-network-tier resources",
						Commands: []*cli.Command{
							{
								Name:  "set-default-network-tier",
								Usage: "set-default-network-tier set-default-network-tier",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-default-network-tier...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-deletion-protection",
						Usage: "Manage set-deletion-protection resources",
						Commands: []*cli.Command{
							{
								Name:  "set-deletion-protection",
								Usage: "set-deletion-protection set-deletion-protection",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-deletion-protection...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-disk-auto-delete",
						Usage: "Manage set-disk-auto-delete resources",
						Commands: []*cli.Command{
							{
								Name:  "set-disk-auto-delete",
								Usage: "set-disk-auto-delete set-disk-auto-delete",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-disk-auto-delete...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-edge-security-policy",
						Usage: "Manage set-edge-security-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "set-edge-security-policy",
								Usage: "set-edge-security-policy set-edge-security-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-edge-security-policy...")
									return nil
								},
							},
							{
								Name:  "set-edge-security-policy",
								Usage: "set-edge-security-policy set-edge-security-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-edge-security-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "Manage set-iam-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy set-iam-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-instance-template",
						Usage: "Manage set-instance-template resources",
						Commands: []*cli.Command{
							{
								Name:  "set-instance-template",
								Usage: "set-instance-template set-instance-template",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-instance-template...")
									return nil
								},
							},
							{
								Name:  "set-instance-template",
								Usage: "set-instance-template set-instance-template",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-instance-template...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-labels",
						Usage: "Manage set-labels resources",
						Commands: []*cli.Command{
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
							{
								Name:  "set-labels",
								Usage: "set-labels set-labels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-labels...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-machine-resources",
						Usage: "Manage set-machine-resources resources",
						Commands: []*cli.Command{
							{
								Name:  "set-machine-resources",
								Usage: "set-machine-resources set-machine-resources",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-machine-resources...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-machine-type",
						Usage: "Manage set-machine-type resources",
						Commands: []*cli.Command{
							{
								Name:  "set-machine-type",
								Usage: "set-machine-type set-machine-type",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-machine-type...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-metadata",
						Usage: "Manage set-metadata resources",
						Commands: []*cli.Command{
							{
								Name:  "set-metadata",
								Usage: "set-metadata set-metadata",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-metadata...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-min-cpu-platform",
						Usage: "Manage set-min-cpu-platform resources",
						Commands: []*cli.Command{
							{
								Name:  "set-min-cpu-platform",
								Usage: "set-min-cpu-platform set-min-cpu-platform",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-min-cpu-platform...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-name",
						Usage: "Manage set-name resources",
						Commands: []*cli.Command{
							{
								Name:  "set-name",
								Usage: "set-name set-name",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-name...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-named-ports",
						Usage: "Manage set-named-ports resources",
						Commands: []*cli.Command{
							{
								Name:  "set-named-ports",
								Usage: "set-named-ports set-named-ports",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-named-ports...")
									return nil
								},
							},
							{
								Name:  "set-named-ports",
								Usage: "set-named-ports set-named-ports",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-named-ports...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-node-template",
						Usage: "Manage set-node-template resources",
						Commands: []*cli.Command{
							{
								Name:  "set-node-template",
								Usage: "set-node-template set-node-template",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-node-template...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-private-ip-google-access",
						Usage: "Manage set-private-ip-google-access resources",
						Commands: []*cli.Command{
							{
								Name:  "set-private-ip-google-access",
								Usage: "set-private-ip-google-access set-private-ip-google-access",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-private-ip-google-access...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-proxy-header",
						Usage: "Manage set-proxy-header resources",
						Commands: []*cli.Command{
							{
								Name:  "set-proxy-header",
								Usage: "set-proxy-header set-proxy-header",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-proxy-header...")
									return nil
								},
							},
							{
								Name:  "set-proxy-header",
								Usage: "set-proxy-header set-proxy-header",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-proxy-header...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-quic-override",
						Usage: "Manage set-quic-override resources",
						Commands: []*cli.Command{
							{
								Name:  "set-quic-override",
								Usage: "set-quic-override set-quic-override",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-quic-override...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-scheduling",
						Usage: "Manage set-scheduling resources",
						Commands: []*cli.Command{
							{
								Name:  "set-scheduling",
								Usage: "set-scheduling set-scheduling",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-scheduling...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-security-policy",
						Usage: "Manage set-security-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "set-security-policy",
								Usage: "set-security-policy set-security-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-security-policy...")
									return nil
								},
							},
							{
								Name:  "set-security-policy",
								Usage: "set-security-policy set-security-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-security-policy...")
									return nil
								},
							},
							{
								Name:  "set-security-policy",
								Usage: "set-security-policy set-security-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-security-policy...")
									return nil
								},
							},
							{
								Name:  "set-security-policy",
								Usage: "set-security-policy set-security-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-security-policy...")
									return nil
								},
							},
							{
								Name:  "set-security-policy",
								Usage: "set-security-policy set-security-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-security-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-service-account",
						Usage: "Manage set-service-account resources",
						Commands: []*cli.Command{
							{
								Name:  "set-service-account",
								Usage: "set-service-account set-service-account",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-service-account...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-shielded-instance-integrity-policy",
						Usage: "Manage set-shielded-instance-integrity-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "set-shielded-instance-integrity-policy",
								Usage: "set-shielded-instance-integrity-policy set-shielded-instance-integrity-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-shielded-instance-integrity-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-ssl-certificates",
						Usage: "Manage set-ssl-certificates resources",
						Commands: []*cli.Command{
							{
								Name:  "set-ssl-certificates",
								Usage: "set-ssl-certificates set-ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-ssl-certificates...")
									return nil
								},
							},
							{
								Name:  "set-ssl-certificates",
								Usage: "set-ssl-certificates set-ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-ssl-certificates...")
									return nil
								},
							},
							{
								Name:  "set-ssl-certificates",
								Usage: "set-ssl-certificates set-ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-ssl-certificates...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-ssl-policy",
						Usage: "Manage set-ssl-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "set-ssl-policy",
								Usage: "set-ssl-policy set-ssl-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-ssl-policy...")
									return nil
								},
							},
							{
								Name:  "set-ssl-policy",
								Usage: "set-ssl-policy set-ssl-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-ssl-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-tags",
						Usage: "Manage set-tags resources",
						Commands: []*cli.Command{
							{
								Name:  "set-tags",
								Usage: "set-tags set-tags",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-tags...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-target",
						Usage: "Manage set-target resources",
						Commands: []*cli.Command{
							{
								Name:  "set-target",
								Usage: "set-target set-target",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-target...")
									return nil
								},
							},
							{
								Name:  "set-target",
								Usage: "set-target set-target",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-target...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-target-pools",
						Usage: "Manage set-target-pools resources",
						Commands: []*cli.Command{
							{
								Name:  "set-target-pools",
								Usage: "set-target-pools set-target-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-target-pools...")
									return nil
								},
							},
							{
								Name:  "set-target-pools",
								Usage: "set-target-pools set-target-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-target-pools...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-url-map",
						Usage: "Manage set-url-map resources",
						Commands: []*cli.Command{
							{
								Name:  "set-url-map",
								Usage: "set-url-map set-url-map",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-url-map...")
									return nil
								},
							},
							{
								Name:  "set-url-map",
								Usage: "set-url-map set-url-map",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-url-map...")
									return nil
								},
							},
							{
								Name:  "set-url-map",
								Usage: "set-url-map set-url-map",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-url-map...")
									return nil
								},
							},
							{
								Name:  "set-url-map",
								Usage: "set-url-map set-url-map",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-url-map...")
									return nil
								},
							},
						},
					},
					{
						Name:  "set-usage-export-bucket",
						Usage: "Manage set-usage-export-bucket resources",
						Commands: []*cli.Command{
							{
								Name:  "set-usage-export-bucket",
								Usage: "set-usage-export-bucket set-usage-export-bucket",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-usage-export-bucket...")
									return nil
								},
							},
						},
					},
					{
						Name:  "simulate-maintenance-event",
						Usage: "Manage simulate-maintenance-event resources",
						Commands: []*cli.Command{
							{
								Name:  "simulate-maintenance-event",
								Usage: "simulate-maintenance-event simulate-maintenance-event",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing simulate-maintenance-event...")
									return nil
								},
							},
							{
								Name:  "simulate-maintenance-event",
								Usage: "simulate-maintenance-event simulate-maintenance-event",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing simulate-maintenance-event...")
									return nil
								},
							},
						},
					},
					{
						Name:  "snapshot-settings",
						Usage: "Manage snapshot-settings resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe snapshot-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch snapshot-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe snapshot-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch snapshot-settings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
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
								Name:  "delete",
								Usage: "delete snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list snapshots",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "ssl-certificates",
						Usage: "Manage ssl-certificates resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list ssl-certificates",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "ssl-policies",
						Usage: "Manage ssl-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch ssl-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "start",
						Usage: "Manage start resources",
						Commands: []*cli.Command{
							{
								Name:  "start",
								Usage: "start start",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start...")
									return nil
								},
							},
						},
					},
					{
						Name:  "start-async-replication",
						Usage: "Manage start-async-replication resources",
						Commands: []*cli.Command{
							{
								Name:  "start-async-replication",
								Usage: "start-async-replication start-async-replication",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-async-replication...")
									return nil
								},
							},
							{
								Name:  "start-async-replication",
								Usage: "start-async-replication start-async-replication",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-async-replication...")
									return nil
								},
							},
						},
					},
					{
						Name:  "start-instances",
						Usage: "Manage start-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "start-instances",
								Usage: "start-instances start-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-instances...")
									return nil
								},
							},
							{
								Name:  "start-instances",
								Usage: "start-instances start-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "start-with-encryption-key",
						Usage: "Manage start-with-encryption-key resources",
						Commands: []*cli.Command{
							{
								Name:  "start-with-encryption-key",
								Usage: "start-with-encryption-key start-with-encryption-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing start-with-encryption-key...")
									return nil
								},
							},
						},
					},
					{
						Name:  "stop",
						Usage: "Manage stop resources",
						Commands: []*cli.Command{
							{
								Name:  "stop",
								Usage: "stop stop",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop...")
									return nil
								},
							},
						},
					},
					{
						Name:  "stop-async-replication",
						Usage: "Manage stop-async-replication resources",
						Commands: []*cli.Command{
							{
								Name:  "stop-async-replication",
								Usage: "stop-async-replication stop-async-replication",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop-async-replication...")
									return nil
								},
							},
							{
								Name:  "stop-async-replication",
								Usage: "stop-async-replication stop-async-replication",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop-async-replication...")
									return nil
								},
							},
						},
					},
					{
						Name:  "stop-group-async-replication",
						Usage: "Manage stop-group-async-replication resources",
						Commands: []*cli.Command{
							{
								Name:  "stop-group-async-replication",
								Usage: "stop-group-async-replication stop-group-async-replication",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop-group-async-replication...")
									return nil
								},
							},
							{
								Name:  "stop-group-async-replication",
								Usage: "stop-group-async-replication stop-group-async-replication",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop-group-async-replication...")
									return nil
								},
							},
						},
					},
					{
						Name:  "stop-instances",
						Usage: "Manage stop-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "stop-instances",
								Usage: "stop-instances stop-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop-instances...")
									return nil
								},
							},
							{
								Name:  "stop-instances",
								Usage: "stop-instances stop-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing stop-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "storage-pool-types",
						Usage: "Manage storage-pool-types resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list storage-pool-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe storage-pool-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list storage-pool-types",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
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
								Name:  "aggregated-list",
								Usage: "aggregated-list storage-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete storage-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe storage-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert storage-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list storage-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update storage-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "subnetworks",
						Usage: "Manage subnetworks resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list subnetworks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete subnetworks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe subnetworks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert subnetworks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list subnetworks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch subnetworks",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "suspend",
						Usage: "Manage suspend resources",
						Commands: []*cli.Command{
							{
								Name:  "suspend",
								Usage: "suspend suspend",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing suspend...")
									return nil
								},
							},
						},
					},
					{
						Name:  "suspend-instances",
						Usage: "Manage suspend-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "suspend-instances",
								Usage: "suspend-instances suspend-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing suspend-instances...")
									return nil
								},
							},
							{
								Name:  "suspend-instances",
								Usage: "suspend-instances suspend-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing suspend-instances...")
									return nil
								},
							},
						},
					},
					{
						Name:  "switch-to-custom-mode",
						Usage: "Manage switch-to-custom-mode resources",
						Commands: []*cli.Command{
							{
								Name:  "switch-to-custom-mode",
								Usage: "switch-to-custom-mode switch-to-custom-mode",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing switch-to-custom-mode...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-grpc-proxies",
						Usage: "Manage target-grpc-proxies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete target-grpc-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-grpc-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-grpc-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-grpc-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch target-grpc-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-http-proxies",
						Usage: "Manage target-http-proxies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch target-http-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-https-proxies",
						Usage: "Manage target-https-proxies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch target-https-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-instances",
						Usage: "Manage target-instances resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list target-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete target-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-instances",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-pools",
						Usage: "Manage target-pools resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list target-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete target-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-pools",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-ssl-proxies",
						Usage: "Manage target-ssl-proxies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete target-ssl-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-ssl-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-ssl-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-ssl-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-tcp-proxies",
						Usage: "Manage target-tcp-proxies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-tcp-proxies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "target-vpn-gateways",
						Usage: "Manage target-vpn-gateways resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list target-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete target-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe target-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert target-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list target-vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "Manage test-iam-permissions resources",
						Commands: []*cli.Command{
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions test-iam-permissions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-access-config",
						Usage: "Manage update-access-config resources",
						Commands: []*cli.Command{
							{
								Name:  "update-access-config",
								Usage: "update-access-config update-access-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-access-config...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-display-device",
						Usage: "Manage update-display-device resources",
						Commands: []*cli.Command{
							{
								Name:  "update",
								Usage: "update update-display-device",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-kms-key",
						Usage: "Manage update-kms-key resources",
						Commands: []*cli.Command{
							{
								Name:  "update-kms-key",
								Usage: "update-kms-key update-kms-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-kms-key...")
									return nil
								},
							},
							{
								Name:  "update-kms-key",
								Usage: "update-kms-key update-kms-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-kms-key...")
									return nil
								},
							},
							{
								Name:  "update-kms-key",
								Usage: "update-kms-key update-kms-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-kms-key...")
									return nil
								},
							},
							{
								Name:  "update-kms-key",
								Usage: "update-kms-key update-kms-key",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-kms-key...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-network-interface",
						Usage: "Manage update-network-interface resources",
						Commands: []*cli.Command{
							{
								Name:  "update",
								Usage: "update update-network-interface",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-peering",
						Usage: "Manage update-peering resources",
						Commands: []*cli.Command{
							{
								Name:  "update",
								Usage: "update update-peering",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-per-instance-configs",
						Usage: "Manage update-per-instance-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "update-per-instance-configs",
								Usage: "update-per-instance-configs update-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-per-instance-configs...")
									return nil
								},
							},
							{
								Name:  "update-per-instance-configs",
								Usage: "update-per-instance-configs update-per-instance-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-per-instance-configs...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-route-policy",
						Usage: "Manage update-route-policy resources",
						Commands: []*cli.Command{
							{
								Name:  "update-route-policy",
								Usage: "update-route-policy update-route-policy",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update-route-policy...")
									return nil
								},
							},
						},
					},
					{
						Name:  "update-shielded-instance-config",
						Usage: "Manage update-shielded-instance-config resources",
						Commands: []*cli.Command{
							{
								Name:  "update",
								Usage: "update update-shielded-instance-config",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "url-maps",
						Usage: "Manage url-maps resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update url-maps",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "validate",
						Usage: "Manage validate resources",
						Commands: []*cli.Command{
							{
								Name:  "validate",
								Usage: "validate validate",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing validate...")
									return nil
								},
							},
							{
								Name:  "validate",
								Usage: "validate validate",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing validate...")
									return nil
								},
							},
						},
					},
					{
						Name:  "vm-extension-policies",
						Usage: "Manage vm-extension-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete vm-extension-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe vm-extension-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert vm-extension-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list vm-extension-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update vm-extension-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
						},
					},
					{
						Name:  "vpn-gateways",
						Usage: "Manage vpn-gateways resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list vpn-gateways",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "vpn-tunnels",
						Usage: "Manage vpn-tunnels resources",
						Commands: []*cli.Command{
							{
								Name:  "aggregated-list",
								Usage: "aggregated-list vpn-tunnels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing aggregated-list...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete vpn-tunnels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe vpn-tunnels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert vpn-tunnels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list vpn-tunnels",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "wait",
						Usage: "Manage wait resources",
						Commands: []*cli.Command{
							{
								Name:  "wait",
								Usage: "wait wait",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing wait...")
									return nil
								},
							},
							{
								Name:  "wait",
								Usage: "wait wait",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing wait...")
									return nil
								},
							},
							{
								Name:  "wait",
								Usage: "wait wait",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing wait...")
									return nil
								},
							},
						},
					},
					{
						Name:  "wire-groups",
						Usage: "Manage wire-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "delete",
								Usage: "delete wire-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing delete...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe wire-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "insert",
								Usage: "insert wire-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing insert...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list wire-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "patch",
								Usage: "patch wire-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing patch...")
									return nil
								},
							},
						},
					},
					{
						Name:  "withdraw",
						Usage: "Manage withdraw resources",
						Commands: []*cli.Command{
							{
								Name:  "withdraw",
								Usage: "withdraw withdraw",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing withdraw...")
									return nil
								},
							},
							{
								Name:  "withdraw",
								Usage: "withdraw withdraw",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing withdraw...")
									return nil
								},
							},
						},
					},
					{
						Name:  "zones",
						Usage: "Manage zones resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list zones",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe zones",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list zones",
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
