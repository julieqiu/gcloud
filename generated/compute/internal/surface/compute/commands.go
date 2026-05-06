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

package compute

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the compute command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "compute",
		Usage: "manage Google Compute Engine API resources",
		Commands: []*cli.Command{
			{
				Name:  "abandon-instances",
				Usage: "Manage abandon-instances resources",
				Commands: []*cli.Command{

					{
						Name:  "abandon-instances",
						Usage: "abandon-instances abandon-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing abandon-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "abandon-instances",
						Usage: "abandon-instances abandon-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing abandon-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe accelerator-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "accelerator_type", Usage: "The ID of the accelerator_type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("accelerator_type"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list accelerator-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "network-interface", Usage: "The name of the network interface to add to this instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing add-access-config on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.BoolFlag{Name: "replace-existing-association", Usage: "Indicates whether or not to replace it if an association of the attachment.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing add-association on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "add-association",
						Usage: "add-association add-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "replace-existing-association", Usage: "Indicates whether or not to replace it if an association of the attachment.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing add-association on %s\n", project)
							return nil
						},
					},

					{
						Name:  "add-association",
						Usage: "add-association add-association",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "replace-existing-association", Usage: "Indicates whether or not to replace it if an association of the attachment.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing add-association on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "add-association",
						Usage: "add-association add-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "replace-existing-association", Usage: "Indicates whether or not to replace it if an association already exists.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing add-association on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing add-health-check on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing add-instance on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group"))
							fmt.Printf("Executing add-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing add-network-interface on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing add-nodes on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "max-priority", Usage: "When rule.", Required: false},
							&cli.IntFlag{Name: "min-priority", Usage: "When rule.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing add-packet-mirroring-rule on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing add-peering on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing add-resource-policies on %s\n", project)
							return nil
						},
					},

					{
						Name:  "add-resource-policies",
						Usage: "add-resource-policies add-resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing add-resource-policies on %s\n", project)
							return nil
						},
					},

					{
						Name:  "add-resource-policies",
						Usage: "add-resource-policies add-resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing add-resource-policies on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing add-rule on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "add-rule",
						Usage: "add-rule add-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "max-priority", Usage: "When rule.", Required: false},
							&cli.IntFlag{Name: "min-priority", Usage: "When rule.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing add-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "add-rule",
						Usage: "add-rule add-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing add-rule on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "add-rule",
						Usage: "add-rule add-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "max-priority", Usage: "When rule.", Required: false},
							&cli.IntFlag{Name: "min-priority", Usage: "When rule.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing add-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "add-rule",
						Usage: "add-rule add-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("security_policy"))
							fmt.Printf("Executing add-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "add-rule",
						Usage: "add-rule add-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("security_policy"))
							fmt.Printf("Executing add-rule on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_bucket"))
							fmt.Printf("Executing add-signed-url-key on %s\n", project)
							return nil
						},
					},

					{
						Name:  "add-signed-url-key",
						Usage: "add-signed-url-key add-signed-url-key",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing add-signed-url-key on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address", Usage: "The ID of the address.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("address"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address", Usage: "The ID of the address.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("address"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address", Usage: "The ID of the address.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("address"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address", Usage: "The ID of the address.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("address"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_advertised_prefix", Usage: "The ID of the public_advertised_prefix.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_advertised_prefix"))
							fmt.Printf("Executing announce on %s\n", project)
							return nil
						},
					},

					{
						Name:  "announce",
						Usage: "announce announce",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing announce on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing apply-updates-to-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "apply-updates-to-instances",
						Usage: "apply-updates-to-instances apply-updates-to-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing apply-updates-to-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force-attach", Usage: "Whether to force attach the regional disk even if it's currently attached.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing attach-disk on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing attach-network-endpoints on %s\n", project)
							return nil
						},
					},

					{
						Name:  "attach-network-endpoints",
						Usage: "attach-network-endpoints attach-network-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing attach-network-endpoints on %s\n", project)
							return nil
						},
					},

					{
						Name:  "attach-network-endpoints",
						Usage: "attach-network-endpoints attach-network-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing attach-network-endpoints on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "The ID of the autoscaler.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("autoscaler"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "The ID of the autoscaler.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("autoscaler"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "Name of the autoscaler to patch.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "Name of the autoscaler to update.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "The ID of the autoscaler.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("autoscaler"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "The ID of the autoscaler.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("autoscaler"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "Name of the autoscaler to patch.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update autoscalers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "autoscaler", Usage: "Name of the autoscaler to update.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_bucket"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_bucket"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_bucket"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_bucket"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_bucket"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_bucket"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch backend-buckets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_bucket"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_service"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_service"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_service"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update backend-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_service"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing bulk-insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "bulk-insert",
						Usage: "bulk-insert bulk-insert",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing bulk-insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "bulk-insert",
						Usage: "bulk-insert bulk-insert",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing bulk-insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "bulk-insert",
						Usage: "bulk-insert bulk-insert",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing bulk-insert on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing bulk-set-labels on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing calendar-mode on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "future_reservation", Usage: "The ID of the future_reservation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("future_reservation"))
							fmt.Printf("Executing cancel on %s\n", project)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel cancel",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resize_request", Usage: "The ID of the resize_request.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"), cmd.String("resize_request"))
							fmt.Printf("Executing cancel on %s\n", project)
							return nil
						},
					},

					{
						Name:  "cancel",
						Usage: "cancel cancel",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resize_request", Usage: "The ID of the resize_request.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"), cmd.String("resize_request"))
							fmt.Printf("Executing cancel on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-firewall-policy", Usage: "The firewall policy from which to copy rules.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing clone-rules on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "clone-rules",
						Usage: "clone-rules clone-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-firewall-policy", Usage: "The firewall policy from which to copy rules.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing clone-rules on %s\n", project)
							return nil
						},
					},

					{
						Name:  "clone-rules",
						Usage: "clone-rules clone-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-firewall-policy", Usage: "The firewall policy from which to copy rules.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing clone-rules on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commitment", Usage: "The ID of the commitment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("commitment"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "commitment", Usage: "The ID of the commitment.", Required: true},
							&cli.StringFlag{Name: "paths", Usage: "The paths.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("commitment"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete composite-health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "composite_health_check", Usage: "The ID of the composite_health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("composite_health_check"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe composite-health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "composite_health_check", Usage: "The ID of the composite_health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("composite_health_check"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert composite-health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list composite-health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch composite-health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "composite_health_check", Usage: "The ID of the composite_health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("composite_health_check"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
							&cli.StringFlag{Name: "source-security-policy", Usage: "The security policy from which to copy rules.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing copy-rules on %s\n", security_policy)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing create on %s\n", project)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create create-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing create on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_group", Usage: "The ID of the interconnect_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_group"))
							fmt.Printf("Executing create on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.BoolFlag{Name: "guest-flush", Usage: "[Input Only] Whether to attempt an application consistent snapshot by.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing create on %s\n", project)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create create-snapshot",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing create on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("cross_site_network"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe cross-site-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("cross_site_network"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert cross-site-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "[Input Only] Validate the new configuration, but don't create it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list cross-site-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch cross-site-networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "[Input Only] Validate the new configuration, but don't update it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("cross_site_network"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access-config", Usage: "The name of the access config to delete.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "network-interface", Usage: "The name of the network interface.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing delete-access-config on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing delete-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete-instances",
						Usage: "delete-instances delete-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing delete-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "network-interface-name", Usage: "The name of the dynamic network interface to be deleted from the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing delete-network-interface on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing delete-nodes on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing delete-per-instance-configs on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete-per-instance-configs",
						Usage: "delete-per-instance-configs delete-per-instance-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing delete-per-instance-configs on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy", Usage: "The Policy name for this request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing delete-route-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "key-name", Usage: "The name of the Signed URL Key to delete.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_bucket"))
							fmt.Printf("Executing delete-signed-url-key on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete-signed-url-key",
						Usage: "delete-signed-url-key delete-signed-url-key",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "key-name", Usage: "The name of the Signed URL Key to delete.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing delete-signed-url-key on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image", Usage: "The ID of the image.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("image"))
							fmt.Printf("Executing deprecate on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-name", Usage: "The device name of the disk to detach.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing detach-disk on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing detach-network-endpoints on %s\n", project)
							return nil
						},
					},

					{
						Name:  "detach-network-endpoints",
						Usage: "detach-network-endpoints detach-network-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing detach-network-endpoints on %s\n", project)
							return nil
						},
					},

					{
						Name:  "detach-network-endpoints",
						Usage: "detach-network-endpoints detach-network-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing detach-network-endpoints on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing disable-xpn-host on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing disable-xpn-resource on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe disk-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk_type", Usage: "The ID of the disk_type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk_type"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list disk-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe disk-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk_type", Usage: "The ID of the disk_type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk_type"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list disk-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-image", Usage: "Source image to restore onto a disk.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "paths", Usage: "The paths.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-image", Usage: "Source image to restore onto a disk.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update disks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "paths", Usage: "The paths.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing enable-xpn-host on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing enable-xpn-resource on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "subnetwork", Usage: "The ID of the subnetwork.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("subnetwork"))
							fmt.Printf("Executing expand-ip-cidr-range on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external_vpn_gateway", Usage: "The ID of the external_vpn_gateway.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("external_vpn_gateway"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe external-vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "external_vpn_gateway", Usage: "The ID of the external_vpn_gateway.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("external_vpn_gateway"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert external-vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list external-vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "family", Usage: "The ID of the family.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("family"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing delete on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.InsertFirewallPolicyRequest{
								ParentId:  cmd.String("parent-id"),
								RequestId: runtime.Ptr(cmd.String("request-id")),
							}

							resp, err := client.Insert(ctx, req)
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
						Name:  "list",
						Usage: "list firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.ListFirewallPoliciesRequest{
								Filter:               runtime.Ptr(cmd.String("filter")),
								MaxResults:           runtime.Ptr(uint32(cmd.Int("max-results"))),
								OrderBy:              runtime.Ptr(cmd.String("order-by")),
								PageToken:            runtime.Ptr(cmd.String("page-token")),
								ParentId:             runtime.Ptr(cmd.String("parent-id")),
								ReturnPartialSuccess: runtime.Ptr(cmd.Bool("return-partial-success")),
							}

							limit := cmd.Int("limit")
							it := client.List(ctx, req)
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
						Name:  "patch",
						Usage: "patch firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing patch on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch firewall-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall", Usage: "The ID of the firewall.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe firewalls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall", Usage: "The ID of the firewall.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert firewalls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list firewalls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch firewalls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall", Usage: "The ID of the firewall.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update firewalls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall", Usage: "The ID of the firewall.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch forwarding-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete future-reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "future_reservation", Usage: "The ID of the future_reservation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("future_reservation"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe future-reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "future_reservation", Usage: "The ID of the future_reservation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("future_reservation"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert future-reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list future-reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update future-reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "future_reservation", Usage: "The ID of the future_reservation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("future_reservation"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing describe on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect", Usage: "The ID of the interconnect.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "network-interface", Usage: "The name of the network interface to get the effective firewalls.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-effective-firewalls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-effective-firewalls",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "Network reference.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query-path", Usage: "Specifies the guest attributes path to be queried.", Required: false},
							&cli.StringFlag{Name: "variable-key", Usage: "Specifies the key for the guest attributes entry.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing get-health on %s\n", project)
							return nil
						},
					},

					{
						Name:  "get-health",
						Usage: "get-health get-health",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_service"))
							fmt.Printf("Executing get-health on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-health",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "composite_health_check", Usage: "The ID of the composite_health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("composite_health_check"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-health",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_source", Usage: "The ID of the health_source.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_source"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "get-health",
						Usage: "get-health get-health",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing get-health on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("%s", cmd.String("resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.GetIamPolicyFirewallPolicyRequest{
								Resource:                      resource,
								OptionsRequestedPolicyVersion: runtime.Ptr(int32(cmd.Int("options-requested-policy-version"))),
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "parent_resource", Usage: "The ID of the parent_resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("parent_resource"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "parent-resource", Usage: "Name or id of parent resource of the resource for this request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-iam-policy",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "options-requested-policy-version", Usage: "Requested IAM Policy version.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect", Usage: "The ID of the interconnect.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "nat-name", Usage: "Name of the nat service to filter the NAT IP information.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "nat-name", Usage: "Name of the nat service to filter the Nat Mapping information.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_attachment_group", Usage: "The ID of the interconnect_attachment_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_attachment_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-operational-status",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_group", Usage: "The ID of the interconnect_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_group"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to get from the firewall policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "policy", Usage: "The Policy name for this request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to get from the firewall policy.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to get from the firewall policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to get from the security policy.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing describe on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to get from the firewall policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to get from the security policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("security_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe get-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to get from the security policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("security_policy"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "vpn_gateway", Usage: "The ID of the vpn_gateway.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("vpn_gateway"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation-sub-block", Usage: "The ID of the reservation sub block.", Required: true},
							&cli.StringFlag{Name: "reservation_slot", Usage: "The ID of the reservation_slot.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/reservationSubBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation-sub-block"), cmd.String("reservation_slot"))
							fmt.Printf("Executing get-version on %s\n", project)
							return nil
						},
					},

					{
						Name:  "get-version",
						Usage: "get-version get-version",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation_sub_block", Usage: "The ID of the reservation_sub_block.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation_sub_block"))
							fmt.Printf("Executing get-version on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete health-aggregation-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_aggregation_policy", Usage: "The ID of the health_aggregation_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_aggregation_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe health-aggregation-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_aggregation_policy", Usage: "The ID of the health_aggregation_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_aggregation_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert health-aggregation-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list health-aggregation-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch health-aggregation-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_aggregation_policy", Usage: "The ID of the health_aggregation_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_aggregation_policy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete health-check-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check_service", Usage: "The ID of the health_check_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_check_service"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe health-check-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check_service", Usage: "The ID of the health_check_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_check_service"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert health-check-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list health-check-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch health-check-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check_service", Usage: "The ID of the health_check_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_check_service"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("health_check"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("health_check"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("health_check"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("health_check"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_check"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_check"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_check"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update health-checks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_check", Usage: "The ID of the health_check.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_check"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete health-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_source", Usage: "The ID of the health_source.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_source"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe health-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_source", Usage: "The ID of the health_source.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_source"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert health-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list health-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch health-sources",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "health_source", Usage: "The ID of the health_source.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("health_source"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "family", Usage: "The ID of the family.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("family"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image", Usage: "The ID of the image.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("image"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image", Usage: "The ID of the image.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("image"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert images",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "force-create", Usage: "Force image creation if true.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "image", Usage: "The ID of the image.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("image"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch instance-group-managers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instance-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instance-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instance-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instance-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instance-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instance-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch instance-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_template", Usage: "The ID of the instance_template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance_template"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_template", Usage: "The ID of the instance_template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("instance_template"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_template", Usage: "The ID of the instance_template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_template"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_template", Usage: "The ID of the instance_template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_template"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instance-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-instance-template", Usage: "Specifies instance template to create the instance.", Required: false},
							&cli.StringFlag{Name: "source-machine-image", Usage: "Specifies the machine image to use to create the instance.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update instances",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "discard-local-ssd", Usage: "Whether to discard local SSDs from the instance during restart.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "minimal-action", Usage: "Specifies the action to take when updating an instance even if the.", Required: false},
							&cli.StringFlag{Name: "most-disruptive-allowed-action", Usage: "Specifies the most disruptive action that can be taken on the instance.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot_group", Usage: "The ID of the instant_snapshot_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instant_snapshot_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instant-snapshot-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot_group", Usage: "The ID of the instant_snapshot_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instant_snapshot_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instant-snapshot-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-consistency-group", Usage: "begin_interface: MixerMutationRequestBuilder.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instant-snapshot-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instant-snapshot-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot_group", Usage: "The ID of the instant_snapshot_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instant_snapshot_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instant-snapshot-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot_group", Usage: "The ID of the instant_snapshot_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instant_snapshot_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instant-snapshot-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-consistency-group", Usage: "begin_interface: MixerMutationRequestBuilder.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instant-snapshot-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot", Usage: "The ID of the instant_snapshot.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instant_snapshot"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot", Usage: "The ID of the instant_snapshot.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instant_snapshot"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot", Usage: "The ID of the instant_snapshot.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instant_snapshot"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instant_snapshot", Usage: "The ID of the instant_snapshot.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instant_snapshot"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list instant-snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_attachment_group", Usage: "The ID of the interconnect_attachment_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_attachment_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe interconnect-attachment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_attachment_group", Usage: "The ID of the interconnect_attachment_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_attachment_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert interconnect-attachment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list interconnect-attachment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch interconnect-attachment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_attachment_group", Usage: "The ID of the interconnect_attachment_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "The list of fields to update.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_attachment_group"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_attachment", Usage: "The ID of the interconnect_attachment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("interconnect_attachment"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_attachment", Usage: "The ID of the interconnect_attachment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("interconnect_attachment"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch interconnect-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_attachment", Usage: "The ID of the interconnect_attachment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("interconnect_attachment"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_group", Usage: "The ID of the interconnect_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe interconnect-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_group", Usage: "The ID of the interconnect_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert interconnect-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list interconnect-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch interconnect-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_group", Usage: "The ID of the interconnect_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "The list of fields to update.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_group"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_location", Usage: "The ID of the interconnect_location.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_location"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list interconnect-locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect_remote_location", Usage: "The ID of the interconnect_remote_location.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect_remote_location"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list interconnect-remote-locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect", Usage: "The ID of the interconnect.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect", Usage: "The ID of the interconnect.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch interconnects",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "interconnect", Usage: "The ID of the interconnect.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("interconnect"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("url_map"))
							fmt.Printf("Executing invalidate-cache on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "license_code", Usage: "The ID of the license_code.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("license_code"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "license", Usage: "The ID of the license.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("license"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe licenses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "license", Usage: "The ID of the license.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("license"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert licenses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list licenses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update licenses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "license", Usage: "The ID of the license.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("license"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "include-inherited-policies", Usage: "If set to \"true\", the response will contain a list of all associations for.", Required: false},
							&cli.StringFlag{Name: "target-resource", Usage: "The target resource to list associations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.ListAssociationsFirewallPolicyRequest{
								IncludeInheritedPolicies: runtime.Ptr(cmd.Bool("include-inherited-policies")),
								TargetResource:           runtime.Ptr(cmd.String("target-resource")),
							}

							resp, err := client.ListAssociations(ctx, req)
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
						Name:  "list",
						Usage: "list list-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "target-resource", Usage: "The target resource to list associations.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.ListAssociationsOrganizationSecurityPolicyRequest{
								TargetResource: runtime.Ptr(cmd.String("target-resource")),
							}

							resp, err := client.ListAssociations(ctx, req)
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
				Name:  "list-available-features",
				Usage: "Manage list-available-features resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list list-available-features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list list-available-features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address-family", Usage: "(Required) limit results to this address family (either IPv4 or IPv6).", Required: false},
							&cli.StringFlag{Name: "destination-prefix", Usage: "Limit results to destinations that are subnets of this CIDR range.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "peer", Usage: "(Required) limit results to the BGP peer with the given name.", Required: false},
							&cli.BoolFlag{Name: "policy-applied", Usage: "When true, the method returns post-policy routes.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "route-type", Usage: "(Required) limit results to this type of route (either LEARNED or.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "storage_pool", Usage: "The ID of the storage_pool.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("storage_pool"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list list-errors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group"))
							fmt.Printf("Executing list-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list-instances",
						Usage: "list-instances list-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group"))
							fmt.Printf("Executing list-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list-managed-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list-managed-instances",
						Usage: "list-managed-instances list-managed-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list-managed-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing list-network-endpoints on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list-network-endpoints",
						Usage: "list-network-endpoints list-network-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing list-network-endpoints on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list-network-endpoints",
						Usage: "list-network-endpoints list-network-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing list-network-endpoints on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing list-nodes on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "direction", Usage: "The direction of the exchanged routes.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "peering-name", Usage: "The response will show routes exchanged over the given peering connection.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The region of the request.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list-per-instance-configs on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list-per-instance-configs",
						Usage: "list-per-instance-configs list-per-instance-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list-per-instance-configs on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.ListPreconfiguredExpressionSetsOrganizationSecurityPoliciesRequest{
								Filter:               runtime.Ptr(cmd.String("filter")),
								MaxResults:           runtime.Ptr(uint32(cmd.Int("max-results"))),
								OrderBy:              runtime.Ptr(cmd.String("order-by")),
								PageToken:            runtime.Ptr(cmd.String("page-token")),
								ParentId:             runtime.Ptr(cmd.String("parent-id")),
								ReturnPartialSuccess: runtime.Ptr(cmd.Bool("return-partial-success")),
							}

							resp, err := client.ListPreconfiguredExpressionSets(ctx, req)
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
						Name:  "list",
						Usage: "list list-preconfigured-expression-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list list-usable",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list list-usable",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list list-usable",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list list-usable",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "service-project", Usage: "The project id or project number in which the subnetwork is intended to be.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list-xpn-hosts on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "machine_image", Usage: "The ID of the machine_image.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("machine_image"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe machine-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "machine_image", Usage: "The ID of the machine_image.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("machine_image"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert machine-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "source-instance", Usage: "Source instance that is used to create the machine image from.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list machine-images",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe machine-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "machine_type", Usage: "The ID of the machine_type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("machine_type"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list machine-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address", Usage: "The ID of the address.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("address"))
							fmt.Printf("Executing move on %s\n", project)
							return nil
						},
					},

					{
						Name:  "move",
						Usage: "move move",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "parent-id", Usage: "The new parent of the firewall policy.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing move on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "move",
						Usage: "move move",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "address", Usage: "The ID of the address.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("address"))
							fmt.Printf("Executing move on %s\n", project)
							return nil
						},
					},

					{
						Name:  "move",
						Usage: "move move",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-id", Usage: "The new parent of the security policy.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing move on %s\n", security_policy)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing move-disk on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing move-instance on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete network-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_attachment", Usage: "The ID of the network_attachment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_attachment"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe network-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_attachment", Usage: "The ID of the network_attachment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_attachment"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert network-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list network-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch network-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_attachment", Usage: "The ID of the network_attachment.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_attachment"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete network-edge-security-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_edge_security_service", Usage: "The ID of the network_edge_security_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_edge_security_service"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe network-edge-security-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_edge_security_service", Usage: "The ID of the network_edge_security_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_edge_security_service"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert network-edge-security-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch network-edge-security-services",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_edge_security_service", Usage: "The ID of the network_edge_security_service.", Required: true},
							&cli.StringFlag{Name: "paths", Usage: "The paths.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "Indicates fields to be updated as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_edge_security_service"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_endpoint_group", Usage: "The ID of the network_endpoint_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("network_endpoint_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list network-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network_profile", Usage: "The ID of the network_profile.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network_profile"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list network-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch networks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete node-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe node-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert node-groups",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "initial-node-count", Usage: "Initial count of nodes in the node group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list node-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch node-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete node-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_template", Usage: "The ID of the node_template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("node_template"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe node-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_template", Usage: "The ID of the node_template.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("node_template"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert node-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list node-templates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe node-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_type", Usage: "The ID of the node_type.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_type"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list node-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete notification-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "notification_endpoint", Usage: "The ID of the notification_endpoint.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("notification_endpoint"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe notification-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "notification_endpoint", Usage: "The ID of the notification_endpoint.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("notification_endpoint"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert notification-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list notification-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							operation := fmt.Sprintf("%s", cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", operation)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							operation := fmt.Sprintf("%s", cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", operation)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.ListGlobalOrganizationOperationsRequest{
								Filter:               runtime.Ptr(cmd.String("filter")),
								MaxResults:           runtime.Ptr(uint32(cmd.Int("max-results"))),
								OrderBy:              runtime.Ptr(cmd.String("order-by")),
								PageToken:            runtime.Ptr(cmd.String("page-token")),
								ParentId:             runtime.Ptr(cmd.String("parent-id")),
								ReturnPartialSuccess: runtime.Ptr(cmd.Bool("return-partial-success")),
							}

							limit := cmd.Int("limit")
							it := client.List(ctx, req)
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
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("operation"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete packet-mirrorings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "packet_mirroring", Usage: "The ID of the packet_mirroring.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("packet_mirroring"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe packet-mirrorings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "packet_mirroring", Usage: "The ID of the packet_mirroring.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("packet_mirroring"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert packet-mirrorings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list packet-mirrorings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch packet-mirrorings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "packet_mirroring", Usage: "The ID of the packet_mirroring.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("packet_mirroring"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to patch.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing patch-packet-mirroring-rule on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing patch-per-instance-configs on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch-per-instance-configs",
						Usage: "patch-per-instance-configs patch-per-instance-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing patch-per-instance-configs on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing patch-route-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to patch.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing patch-rule on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "patch-rule",
						Usage: "patch-rule patch-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to patch.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing patch-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch-rule",
						Usage: "patch-rule patch-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to patch.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing patch-rule on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "patch-rule",
						Usage: "patch-rule patch-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to patch.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing patch-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch-rule",
						Usage: "patch-rule patch-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to patch.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "Indicates fields to be cleared as part of this request.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("security_policy"))
							fmt.Printf("Executing patch-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch-rule",
						Usage: "patch-rule patch-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to patch.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "Indicates fields to be cleared as part of this request.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("security_policy"))
							fmt.Printf("Executing patch-rule on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing perform-maintenance on %s\n", project)
							return nil
						},
					},

					{
						Name:  "perform-maintenance",
						Usage: "perform-maintenance perform-maintenance",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing perform-maintenance on %s\n", project)
							return nil
						},
					},

					{
						Name:  "perform-maintenance",
						Usage: "perform-maintenance perform-maintenance",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation_block", Usage: "The ID of the reservation_block.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation_block"))
							fmt.Printf("Executing perform-maintenance on %s\n", project)
							return nil
						},
					},

					{
						Name:  "perform-maintenance",
						Usage: "perform-maintenance perform-maintenance",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation_sub_block", Usage: "The ID of the reservation_sub_block.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation_sub_block"))
							fmt.Printf("Executing perform-maintenance on %s\n", project)
							return nil
						},
					},

					{
						Name:  "perform-maintenance",
						Usage: "perform-maintenance perform-maintenance",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"))
							fmt.Printf("Executing perform-maintenance on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing preview on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "preview_feature", Usage: "The ID of the preview_feature.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("preview_feature"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list preview-features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update preview-features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "preview_feature", Usage: "The ID of the preview_feature.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("preview_feature"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_advertised_prefix", Usage: "The ID of the public_advertised_prefix.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_advertised_prefix"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe public-advertised-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_advertised_prefix", Usage: "The ID of the public_advertised_prefix.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_advertised_prefix"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert public-advertised-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list public-advertised-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch public-advertised-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_advertised_prefix", Usage: "The ID of the public_advertised_prefix.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_advertised_prefix"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch public-delegated-prefixes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing recreate-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "recreate-instances",
						Usage: "recreate-instances recreate-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing recreate-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list regions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing remove-association on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "remove-association",
						Usage: "remove-association remove-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing remove-association on %s\n", project)
							return nil
						},
					},

					{
						Name:  "remove-association",
						Usage: "remove-association remove-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing remove-association on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "remove-association",
						Usage: "remove-association remove-association",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing remove-association on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing remove-health-check on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing remove-instance on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group"))
							fmt.Printf("Executing remove-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to remove from the firewall policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing remove-packet-mirroring-rule on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing remove-peering on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing remove-resource-policies on %s\n", project)
							return nil
						},
					},

					{
						Name:  "remove-resource-policies",
						Usage: "remove-resource-policies remove-resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing remove-resource-policies on %s\n", project)
							return nil
						},
					},

					{
						Name:  "remove-resource-policies",
						Usage: "remove-resource-policies remove-resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing remove-resource-policies on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to remove from the firewall policy.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							firewall_policy := fmt.Sprintf("%s", cmd.String("firewall_policy"))
							fmt.Printf("Executing remove-rule on %s\n", firewall_policy)
							return nil
						},
					},

					{
						Name:  "remove-rule",
						Usage: "remove-rule remove-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to remove from the firewall policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("firewall_policy"))
							fmt.Printf("Executing remove-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "remove-rule",
						Usage: "remove-rule remove-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to remove from the security policy.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing remove-rule on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "remove-rule",
						Usage: "remove-rule remove-rule",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewall_policy", Usage: "The ID of the firewall_policy.", Required: true},
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to remove from the firewall policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("firewall_policy"))
							fmt.Printf("Executing remove-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "remove-rule",
						Usage: "remove-rule remove-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to remove from the security policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("security_policy"))
							fmt.Printf("Executing remove-rule on %s\n", project)
							return nil
						},
					},

					{
						Name:  "remove-rule",
						Usage: "remove-rule remove-rule",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "priority", Usage: "The priority of the rule to remove from the security policy.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("security_policy"))
							fmt.Printf("Executing remove-rule on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation_sub_block", Usage: "The ID of the reservation_sub_block.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation_sub_block"))
							fmt.Printf("Executing report-faulty on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing report-host-as-faulty on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing request-remove-peering on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation_block", Usage: "The ID of the reservation_block.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "View of the Block.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation_block"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list reservation-blocks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block, formatted as.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation-sub-block", Usage: "The ID of the reservation sub block.", Required: true},
							&cli.StringFlag{Name: "reservation_slot", Usage: "The ID of the reservation_slot.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/reservationSubBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation-sub-block"), cmd.String("reservation_slot"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list reservation-slots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block, formatted as.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation-sub-block", Usage: "The ID of the reservation sub block.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/reservationSubBlocks/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation-sub-block"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update reservation-slots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the sub-block resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation-sub-block", Usage: "The ID of the reservation sub block.", Required: true},
							&cli.StringFlag{Name: "reservation_slot", Usage: "The ID of the reservation_slot.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "The fields to be updated as part of this request.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/reservationSubBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation-sub-block"), cmd.String("reservation_slot"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "reservation_sub_block", Usage: "The ID of the reservation_sub_block.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "View of the subBlock.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("reservation_sub_block"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list reservation-sub-blocks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "parent-name", Usage: "The name of the parent reservation and parent block.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "paths", Usage: "The paths.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "Update_mask indicates fields to be updated as part of this request.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing reset on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing resize on %s\n", project)
							return nil
						},
					},

					{
						Name:  "resize",
						Usage: "resize resize",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.IntFlag{Name: "size", Usage: "The number of running instances that the managed instance group should.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing resize on %s\n", project)
							return nil
						},
					},

					{
						Name:  "resize",
						Usage: "resize resize",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing resize on %s\n", project)
							return nil
						},
					},

					{
						Name:  "resize",
						Usage: "resize resize",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.IntFlag{Name: "size", Usage: "Number of instances that should exist in this instance group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing resize on %s\n", project)
							return nil
						},
					},

					{
						Name:  "resize",
						Usage: "resize resize",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"))
							fmt.Printf("Executing resize on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resize_request", Usage: "The ID of the resize_request.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"), cmd.String("resize_request"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe resize-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resize_request", Usage: "The ID of the resize_request.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"), cmd.String("resize_request"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert resize-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list resize-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete resize-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resize_request", Usage: "The ID of the resize_request.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"), cmd.String("resize_request"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe resize-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resize_request", Usage: "The ID of the resize_request.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"), cmd.String("resize_request"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert resize-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list resize-requests",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource_policy", Usage: "The ID of the resource_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource_policy", Usage: "The ID of the resource_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch resource-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource_policy", Usage: "The ID of the resource_policy.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource_policy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing resume on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing resume-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "resume-instances",
						Usage: "resume-instances resume-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing resume-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update routers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "route", Usage: "The ID of the route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("route"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "route", Usage: "The ID of the route.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("route"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list routes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing delete on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing describe on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: false},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.InsertOrganizationSecurityPolicyRequest{
								ParentId:  runtime.Ptr(cmd.String("parent-id")),
								RequestId: runtime.Ptr(cmd.String("request-id")),
							}

							resp, err := client.Insert(ctx, req)
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
						Name:  "list",
						Usage: "list security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "parent-id", Usage: "Parent ID for this request.", Required: false},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.ListOrganizationSecurityPoliciesRequest{
								Filter:               runtime.Ptr(cmd.String("filter")),
								MaxResults:           runtime.Ptr(uint32(cmd.Int("max-results"))),
								OrderBy:              runtime.Ptr(cmd.String("order-by")),
								PageToken:            runtime.Ptr(cmd.String("page-token")),
								ParentId:             runtime.Ptr(cmd.String("parent-id")),
								ReturnPartialSuccess: runtime.Ptr(cmd.Bool("return-partial-success")),
							}

							limit := cmd.Int("limit")
							it := client.List(ctx, req)
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
						Name:  "patch",
						Usage: "patch security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							security_policy := fmt.Sprintf("%s", cmd.String("security_policy"))
							fmt.Printf("Executing patch on %s\n", security_policy)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("security_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("security_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "Indicates fields to be cleared as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("security_policy"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("security_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("security_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If true, the request will not be committed.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "security_policy", Usage: "The ID of the security_policy.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "Indicates fields to be cleared as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("security_policy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing send-diagnostic-interrupt on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.IntFlag{Name: "port", Usage: "Specifies which COM or serial port to retrieve data from.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "start", Usage: "Specifies the starting byte position of the output to return.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing describe on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete service-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service_attachment", Usage: "The ID of the service_attachment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("service_attachment"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe service-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "service_attachment", Usage: "The ID of the service_attachment.", Required: true},
							&cli.BoolFlag{Name: "show-nat-ips", Usage: "Indicates whether NAT IPs should be included in the response.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("service_attachment"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert service-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list service-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch service-attachments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "service_attachment", Usage: "The ID of the service_attachment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("service_attachment"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_ssl_proxy", Usage: "The ID of the target_ssl_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_ssl_proxy"))
							fmt.Printf("Executing set-backend-service on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-backend-service",
						Usage: "set-backend-service set-backend-service",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_tcp_proxy", Usage: "The ID of the target_tcp_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_tcp_proxy"))
							fmt.Printf("Executing set-backend-service on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.FloatFlag{Name: "failover-ratio", Usage: "New failoverRatio value for the target pool.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing set-backup on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing set-certificate-map on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-certificate-map",
						Usage: "set-certificate-map set-certificate-map",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_ssl_proxy", Usage: "The ID of the target_ssl_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_ssl_proxy"))
							fmt.Printf("Executing set-certificate-map on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing set-cloud-armor-tier on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing set-common-instance-metadata on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing set-default-network-tier on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "deletion-protection", Usage: "Whether the resource should be protected against deletion.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-deletion-protection on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "auto-delete", Usage: "Whether to auto-delete the disk when the instance is deleted.", Required: true},
							&cli.StringFlag{Name: "device-name", Usage: "The device name of the disk to modify.", Required: true},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-disk-auto-delete on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_bucket", Usage: "The ID of the backend_bucket.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_bucket"))
							fmt.Printf("Executing set-edge-security-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-edge-security-policy",
						Usage: "set-edge-security-policy set-edge-security-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing set-edge-security-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("%s", cmd.String("resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.SetIamPolicyFirewallPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent_resource", Usage: "The ID of the parent_resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("parent_resource"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-resource", Usage: "Name or id of parent resource of the resource for this request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy set-iam-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-iam-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing set-instance-template on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-instance-template",
						Usage: "set-instance-template set-instance-template",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing set-instance-template on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-labels",
						Usage: "set-labels set-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing set-labels on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-machine-resources on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-machine-type on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-metadata on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-min-cpu-platform on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-name on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group"))
							fmt.Printf("Executing set-named-ports on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-named-ports",
						Usage: "set-named-ports set-named-ports",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group", Usage: "The ID of the instance_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group"))
							fmt.Printf("Executing set-named-ports on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing set-node-template on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "subnetwork", Usage: "The ID of the subnetwork.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("subnetwork"))
							fmt.Printf("Executing set-private-ip-google-access on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_ssl_proxy", Usage: "The ID of the target_ssl_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_ssl_proxy"))
							fmt.Printf("Executing set-proxy-header on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-proxy-header",
						Usage: "set-proxy-header set-proxy-header",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_tcp_proxy", Usage: "The ID of the target_tcp_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_tcp_proxy"))
							fmt.Printf("Executing set-proxy-header on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing set-quic-override on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-scheduling on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("backend_service"))
							fmt.Printf("Executing set-security-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-security-policy",
						Usage: "set-security-policy set-security-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-security-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-security-policy",
						Usage: "set-security-policy set-security-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "backend_service", Usage: "The ID of the backend_service.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("backend_service"))
							fmt.Printf("Executing set-security-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-security-policy",
						Usage: "set-security-policy set-security-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_instance", Usage: "The ID of the target_instance.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("target_instance"))
							fmt.Printf("Executing set-security-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-security-policy",
						Usage: "set-security-policy set-security-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing set-security-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-service-account on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-shielded-instance-integrity-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing set-ssl-certificates on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-ssl-certificates",
						Usage: "set-ssl-certificates set-ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing set-ssl-certificates on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-ssl-certificates",
						Usage: "set-ssl-certificates set-ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_ssl_proxy", Usage: "The ID of the target_ssl_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_ssl_proxy"))
							fmt.Printf("Executing set-ssl-certificates on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing set-ssl-policy on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-ssl-policy",
						Usage: "set-ssl-policy set-ssl-policy",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_ssl_proxy", Usage: "The ID of the target_ssl_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_ssl_proxy"))
							fmt.Printf("Executing set-ssl-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing set-tags on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing set-target on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-target",
						Usage: "set-target set-target",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "forwarding_rule", Usage: "The ID of the forwarding_rule.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("forwarding_rule"))
							fmt.Printf("Executing set-target on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing set-target-pools on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-target-pools",
						Usage: "set-target-pools set-target-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing set-target-pools on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_http_proxy", Usage: "The ID of the target_http_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_http_proxy"))
							fmt.Printf("Executing set-url-map on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-url-map",
						Usage: "set-url-map set-url-map",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing set-url-map on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-url-map",
						Usage: "set-url-map set-url-map",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_http_proxy", Usage: "The ID of the target_http_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_http_proxy"))
							fmt.Printf("Executing set-url-map on %s\n", project)
							return nil
						},
					},

					{
						Name:  "set-url-map",
						Usage: "set-url-map set-url-map",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing set-url-map on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing set-usage-export-bucket on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "with-extended-notifications", Usage: "Determines whether the customers receive notifications before migration.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing simulate-maintenance-event on %s\n", project)
							return nil
						},
					},

					{
						Name:  "simulate-maintenance-event",
						Usage: "simulate-maintenance-event simulate-maintenance-event",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "node_group", Usage: "The ID of the node_group.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("node_group"))
							fmt.Printf("Executing simulate-maintenance-event on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch snapshot-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe snapshot-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch snapshot-settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("snapshot"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("snapshot"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("snapshot"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("snapshot"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list snapshots",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "ssl_certificate", Usage: "The ID of the ssl_certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("ssl_certificate"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "ssl_certificate", Usage: "The ID of the ssl_certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("ssl_certificate"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "ssl_certificate", Usage: "The ID of the ssl_certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("ssl_certificate"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "ssl_certificate", Usage: "The ID of the ssl_certificate.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("ssl_certificate"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list ssl-certificates",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "ssl_policy", Usage: "The ID of the ssl_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("ssl_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "ssl_policy", Usage: "The ID of the ssl_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("ssl_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "ssl_policy", Usage: "The ID of the ssl_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("ssl_policy"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "ssl_policy", Usage: "The ID of the ssl_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("ssl_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "ssl_policy", Usage: "The ID of the ssl_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("ssl_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch ssl-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "ssl_policy", Usage: "The ID of the ssl_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("ssl_policy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing start on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing start-async-replication on %s\n", project)
							return nil
						},
					},

					{
						Name:  "start-async-replication",
						Usage: "start-async-replication start-async-replication",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing start-async-replication on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing start-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "start-instances",
						Usage: "start-instances start-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing start-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing start-with-encryption-key on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "discard-local-ssd", Usage: "This property is required if the instance has any attached Local SSD disks.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing stop on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing stop-async-replication on %s\n", project)
							return nil
						},
					},

					{
						Name:  "stop-async-replication",
						Usage: "stop-async-replication stop-async-replication",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing stop-async-replication on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing stop-group-async-replication on %s\n", project)
							return nil
						},
					},

					{
						Name:  "stop-group-async-replication",
						Usage: "stop-group-async-replication stop-group-async-replication",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing stop-group-async-replication on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing stop-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "stop-instances",
						Usage: "stop-instances stop-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing stop-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe storage-pool-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage_pool_type", Usage: "The ID of the storage_pool_type.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("storage_pool_type"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list storage-pool-types",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "storage_pool", Usage: "The ID of the storage_pool.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("storage_pool"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "storage_pool", Usage: "The ID of the storage_pool.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("storage_pool"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update storage-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "storage_pool", Usage: "The ID of the storage_pool.", Required: true},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("storage_pool"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
							&cli.StringFlag{Name: "views", Usage: "Defines the extra views returned back in the subnetwork resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete subnetworks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "subnetwork", Usage: "The ID of the subnetwork.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("subnetwork"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe subnetworks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "subnetwork", Usage: "The ID of the subnetwork.", Required: true},
							&cli.StringFlag{Name: "views", Usage: "Defines the extra views returned back in the subnetwork resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("subnetwork"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert subnetworks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list subnetworks",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "views", Usage: "Defines the extra views returned back in the subnetwork resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch subnetworks",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "drain-timeout-seconds", Usage: "The drain timeout specifies the upper bound in seconds on the amount of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "subnetwork", Usage: "The ID of the subnetwork.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("subnetwork"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "discard-local-ssd", Usage: "This property is required if the instance has any attached Local SSD disks.", Required: false},
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing suspend on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing suspend-instances on %s\n", project)
							return nil
						},
					},

					{
						Name:  "suspend-instances",
						Usage: "suspend-instances suspend-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing suspend-instances on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing switch-to-custom-mode on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_grpc_proxy", Usage: "The ID of the target_grpc_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_grpc_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-grpc-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target_grpc_proxy", Usage: "The ID of the target_grpc_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_grpc_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-grpc-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-grpc-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch target-grpc-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_grpc_proxy", Usage: "The ID of the target_grpc_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_grpc_proxy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_http_proxy", Usage: "The ID of the target_http_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_http_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "target_http_proxy", Usage: "The ID of the target_http_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_http_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_http_proxy", Usage: "The ID of the target_http_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_http_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target_http_proxy", Usage: "The ID of the target_http_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_http_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch target-http-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_http_proxy", Usage: "The ID of the target_http_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_http_proxy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch target-https-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_https_proxy", Usage: "The ID of the target_https_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_https_proxy"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_instance", Usage: "The ID of the target_instance.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("target_instance"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target_instance", Usage: "The ID of the target_instance.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("target_instance"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-instances",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "target_pool", Usage: "The ID of the target_pool.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_pool"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-pools",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_ssl_proxy", Usage: "The ID of the target_ssl_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_ssl_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-ssl-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target_ssl_proxy", Usage: "The ID of the target_ssl_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_ssl_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-ssl-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-ssl-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_tcp_proxy", Usage: "The ID of the target_tcp_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_tcp_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "target_tcp_proxy", Usage: "The ID of the target_tcp_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_tcp_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_tcp_proxy", Usage: "The ID of the target_tcp_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_tcp_proxy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "target_tcp_proxy", Usage: "The ID of the target_tcp_proxy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("target_tcp_proxy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-tcp-proxies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete target-vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "target_vpn_gateway", Usage: "The ID of the target_vpn_gateway.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_vpn_gateway"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe target-vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "target_vpn_gateway", Usage: "The ID of the target_vpn_gateway.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("target_vpn_gateway"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert target-vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list target-vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("%s", cmd.String("resource"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := compute.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &computepb.TestIamPermissionsFirewallPolicyRequest{
								Resource: resource,
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent_resource", Usage: "The ID of the parent_resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("parent_resource"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "parent-resource", Usage: "Name or id of parent resource of the resource for this request.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
							&cli.StringFlag{Name: "reservation-block", Usage: "The ID of the reservation block.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/reservations/%s/reservationBlocks/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("reservation"), cmd.String("reservation-block"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
							return nil
						},
					},

					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions test-iam-permissions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "resource", Usage: "The ID of the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("resource"))
							fmt.Printf("Executing test-iam-permissions on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "network-interface", Usage: "The name of the network interface where the access config is attached.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing update-access-config on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("disk"))
							fmt.Printf("Executing update-kms-key on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update-kms-key",
						Usage: "update-kms-key update-kms-key",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "disk", Usage: "The ID of the disk.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("disk"))
							fmt.Printf("Executing update-kms-key on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update-kms-key",
						Usage: "update-kms-key update-kms-key",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("snapshot"))
							fmt.Printf("Executing update-kms-key on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update-kms-key",
						Usage: "update-kms-key update-kms-key",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "snapshot", Usage: "The ID of the snapshot.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("snapshot"))
							fmt.Printf("Executing update-kms-key on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "network-interface", Usage: "The name of the network interface to update.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "network", Usage: "The ID of the network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("network"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing update-per-instance-configs on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update-per-instance-configs",
						Usage: "update-per-instance-configs update-per-instance-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance_group_manager", Usage: "The ID of the instance_group_manager.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("instance_group_manager"))
							fmt.Printf("Executing update-per-instance-configs on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "router", Usage: "The ID of the router.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("router"))
							fmt.Printf("Executing update-route-policy on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "instance", Usage: "The ID of the instance.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("instance"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "begin_interface: MixerMutationRequestBuilder.", Required: false},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("url_map"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("url_map"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "begin_interface: MixerMutationRequestBuilder.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "begin_interface: MixerMutationRequestBuilder.", Required: false},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("url_map"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "begin_interface: MixerMutationRequestBuilder.", Required: false},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("url_map"))
							fmt.Printf("Executing update on %s\n", project)
							return nil
						},
					},

					{
						Name:  "aggregated-list",
						Usage: "aggregated-list url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("url_map"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("url_map"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("url_map"))
							fmt.Printf("Executing patch on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update url-maps",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("url_map"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("url_map"))
							fmt.Printf("Executing validate on %s\n", project)
							return nil
						},
					},

					{
						Name:  "validate",
						Usage: "validate validate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "url_map", Usage: "The ID of the url_map.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("url_map"))
							fmt.Printf("Executing validate on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "vm_extension_policy", Usage: "The ID of the vm_extension_policy.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("vm_extension_policy"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe vm-extension-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "vm_extension_policy", Usage: "The ID of the vm_extension_policy.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("vm_extension_policy"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert vm-extension-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list vm-extension-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update vm-extension-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "vm_extension_policy", Usage: "The ID of the vm_extension_policy.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("vm_extension_policy"))
							fmt.Printf("Executing update on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "vpn_gateway", Usage: "The ID of the vpn_gateway.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("vpn_gateway"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "vpn_gateway", Usage: "The ID of the vpn_gateway.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("vpn_gateway"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list vpn-gateways",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.BoolFlag{Name: "include-all-scopes", Usage: "Indicates whether every visible scope for each scope type (zone, region,.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
							&cli.IntFlag{Name: "service-project-number", Usage: "The Shared VPC service project id or service project number for which.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing aggregated-list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete vpn-tunnels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "vpn_tunnel", Usage: "The ID of the vpn_tunnel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("vpn_tunnel"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe vpn-tunnels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "vpn_tunnel", Usage: "The ID of the vpn_tunnel.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("vpn_tunnel"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert vpn-tunnels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list vpn-tunnels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", project)
							return nil
						},
					},

					{
						Name:  "wait",
						Usage: "wait wait",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", project)
							return nil
						},
					},

					{
						Name:  "wait",
						Usage: "wait wait",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("zone"), cmd.String("operation"))
							fmt.Printf("Executing wait on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "wire_group", Usage: "The ID of the wire_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("cross_site_network"), cmd.String("wire_group"))
							fmt.Printf("Executing delete on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe wire-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "wire_group", Usage: "The ID of the wire_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("cross_site_network"), cmd.String("wire_group"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "insert",
						Usage: "insert wire-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "[Input Only] Validate the new configuration, but don't create it.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("cross_site_network"))
							fmt.Printf("Executing insert on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list wire-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("cross_site_network"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "patch",
						Usage: "patch wire-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "cross_site_network", Usage: "The ID of the cross_site_network.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
							&cli.StringFlag{Name: "update-mask", Usage: "update_mask indicates fields to be updated as part of this request.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "[Input Only] Validate the new configuration, but don't update it.", Required: false},
							&cli.StringFlag{Name: "wire_group", Usage: "The ID of the wire_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("cross_site_network"), cmd.String("wire_group"))
							fmt.Printf("Executing patch on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_advertised_prefix", Usage: "The ID of the public_advertised_prefix.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("public_advertised_prefix"))
							fmt.Printf("Executing withdraw on %s\n", project)
							return nil
						},
					},

					{
						Name:  "withdraw",
						Usage: "withdraw withdraw",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "public_delegated_prefix", Usage: "The ID of the public_delegated_prefix.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.StringFlag{Name: "request-id", Usage: "An optional request ID to identify requests.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s/%s", cmd.String("project"), cmd.String("region"), cmd.String("public_delegated_prefix"))
							fmt.Printf("Executing withdraw on %s\n", project)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "region", Usage: "The Cloud region for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("region"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "zone", Usage: "The Cloud zone for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s/%s", cmd.String("project"), cmd.String("zone"))
							fmt.Printf("Executing describe on %s\n", project)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list zones",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "A filter expression that filters resources listed in the response.", Required: false},
							&cli.IntFlag{Name: "max-results", Usage: "The maximum number of results per page that should be returned.", Required: false},
							&cli.StringFlag{Name: "order-by", Usage: "Sorts list results by a certain order.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Specifies a page token to use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "Opt-in for partial success behavior which provides partial results in case.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project := fmt.Sprintf("%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", project)
							return nil
						},
					},
				},
			},
		},
	}
}
