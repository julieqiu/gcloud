package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	networkservices "cloud.google.com/go/networkservices/apiv1"
	"cloud.google.com/go/networkservices/apiv1/networkservicespb"
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
				Name:  "networkservices",
				Usage: "manage Network Services API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "authz-extensions",
						Usage: "Manage authz-extensions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list authz-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListAuthzExtensionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAuthzExtensions(ctx, req)
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
								Usage: "describe authz-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz_extension", Usage: "The authz_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetAuthzExtensionRequest{Name: name}
									resp, err := client.GetAuthzExtension(ctx, req)
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
								Usage: "create authz-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz-extension-id", Usage: "The authz extension id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "authority", Usage: "The authority.", Required: true},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: true},
									&cli.BoolFlag{Name: "fail-open", Usage: "The fail open.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateAuthzExtensionRequest{Parent: parent}
									req.AuthzExtensionId = cmd.String("authz-extension-id")
									req.AuthzExtension = &networkservicespb.AuthzExtension{
										Description: cmd.String("description"),
										Authority:   cmd.String("authority"),
										Service:     cmd.String("service"),
										FailOpen:    cmd.Bool("fail-open"),
									}
									op, err := client.CreateAuthzExtension(ctx, req)
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
								Usage: "update authz-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz_extension", Usage: "The authz_extension.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "authority", Usage: "The authority.", Required: false},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: false},
									&cli.BoolFlag{Name: "fail-open", Usage: "The fail open.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateAuthzExtensionRequest{}
									req.AuthzExtension = &networkservicespb.AuthzExtension{
										Name:        name,
										Description: cmd.String("description"),
										Authority:   cmd.String("authority"),
										Service:     cmd.String("service"),
										FailOpen:    cmd.Bool("fail-open"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("authority") {
										paths = append(paths, "authority")
									}
									if cmd.IsSet("service") {
										paths = append(paths, "service")
									}
									if cmd.IsSet("fail-open") {
										paths = append(paths, "fail_open")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAuthzExtension(ctx, req)
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
								Usage: "delete authz-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz_extension", Usage: "The authz_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteAuthzExtensionRequest{Name: name}
									op, err := client.DeleteAuthzExtension(ctx, req)
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
						Name:  "edge-cache-keysets",
						Usage: "Manage edge-cache-keysets resources",
						Commands: []*cli.Command{
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy edge-cache-keysets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy edge-cache-keysets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions edge-cache-keysets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy edge-cache-keysets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy edge-cache-keysets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions edge-cache-keysets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "endpoint-policies",
						Usage: "Manage endpoint-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list endpoint-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListEndpointPoliciesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListEndpointPolicies(ctx, req)
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
								Usage: "describe endpoint-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "endpoint_policy", Usage: "The endpoint_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/endpointPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint_policy"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetEndpointPolicyRequest{Name: name}
									resp, err := client.GetEndpointPolicy(ctx, req)
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
								Usage: "create endpoint-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "endpoint-policy-id", Usage: "The endpoint policy id.", Required: true},
									&cli.StringFlag{Name: "authorization-policy", Usage: "The authorization policy.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "server-tls-policy", Usage: "The server tls policy.", Required: false},
									&cli.StringFlag{Name: "client-tls-policy", Usage: "The client tls policy.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateEndpointPolicyRequest{Parent: parent}
									req.EndpointPolicyId = cmd.String("endpoint-policy-id")
									req.EndpointPolicy = &networkservicespb.EndpointPolicy{
										AuthorizationPolicy: cmd.String("authorization-policy"),
										Description:         cmd.String("description"),
										ServerTlsPolicy:     cmd.String("server-tls-policy"),
										ClientTlsPolicy:     cmd.String("client-tls-policy"),
									}
									op, err := client.CreateEndpointPolicy(ctx, req)
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
								Usage: "update endpoint-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "endpoint_policy", Usage: "The endpoint_policy.", Required: true},
									&cli.StringFlag{Name: "authorization-policy", Usage: "The authorization policy.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "server-tls-policy", Usage: "The server tls policy.", Required: false},
									&cli.StringFlag{Name: "client-tls-policy", Usage: "The client tls policy.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/endpointPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint_policy"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateEndpointPolicyRequest{}
									req.EndpointPolicy = &networkservicespb.EndpointPolicy{
										Name:                name,
										AuthorizationPolicy: cmd.String("authorization-policy"),
										Description:         cmd.String("description"),
										ServerTlsPolicy:     cmd.String("server-tls-policy"),
										ClientTlsPolicy:     cmd.String("client-tls-policy"),
									}
									var paths []string
									if cmd.IsSet("authorization-policy") {
										paths = append(paths, "authorization_policy")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("server-tls-policy") {
										paths = append(paths, "server_tls_policy")
									}
									if cmd.IsSet("client-tls-policy") {
										paths = append(paths, "client_tls_policy")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateEndpointPolicy(ctx, req)
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
								Usage: "delete endpoint-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "endpoint_policy", Usage: "The endpoint_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/endpointPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("endpoint_policy"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteEndpointPolicyRequest{Name: name}
									op, err := client.DeleteEndpointPolicy(ctx, req)
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
						Name:  "gateways",
						Usage: "Manage gateways resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListGatewaysRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListGateways(ctx, req)
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
								Usage: "describe gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetGatewayRequest{Name: name}
									resp, err := client.GetGateway(ctx, req)
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
								Usage: "create gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway-id", Usage: "The gateway id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: false},
									&cli.StringFlag{Name: "server-tls-policy", Usage: "The server tls policy.", Required: false},
									&cli.StringFlag{Name: "gateway-security-policy", Usage: "The gateway security policy.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.StringFlag{Name: "subnetwork", Usage: "The subnetwork.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateGatewayRequest{Parent: parent}
									req.GatewayId = cmd.String("gateway-id")
									req.Gateway = &networkservicespb.Gateway{
										Description:           cmd.String("description"),
										Scope:                 cmd.String("scope"),
										ServerTlsPolicy:       cmd.String("server-tls-policy"),
										GatewaySecurityPolicy: cmd.String("gateway-security-policy"),
										Network:               cmd.String("network"),
										Subnetwork:            cmd.String("subnetwork"),
									}
									op, err := client.CreateGateway(ctx, req)
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
								Usage: "update gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: false},
									&cli.StringFlag{Name: "server-tls-policy", Usage: "The server tls policy.", Required: false},
									&cli.StringFlag{Name: "gateway-security-policy", Usage: "The gateway security policy.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.StringFlag{Name: "subnetwork", Usage: "The subnetwork.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateGatewayRequest{}
									req.Gateway = &networkservicespb.Gateway{
										Name:                  name,
										Description:           cmd.String("description"),
										Scope:                 cmd.String("scope"),
										ServerTlsPolicy:       cmd.String("server-tls-policy"),
										GatewaySecurityPolicy: cmd.String("gateway-security-policy"),
										Network:               cmd.String("network"),
										Subnetwork:            cmd.String("subnetwork"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("scope") {
										paths = append(paths, "scope")
									}
									if cmd.IsSet("server-tls-policy") {
										paths = append(paths, "server_tls_policy")
									}
									if cmd.IsSet("gateway-security-policy") {
										paths = append(paths, "gateway_security_policy")
									}
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									if cmd.IsSet("subnetwork") {
										paths = append(paths, "subnetwork")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateGateway(ctx, req)
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
								Usage: "delete gateways",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteGatewayRequest{Name: name}
									op, err := client.DeleteGateway(ctx, req)
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
						Name:  "grpc-routes",
						Usage: "Manage grpc-routes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list grpc-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListGrpcRoutesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListGrpcRoutes(ctx, req)
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
								Usage: "describe grpc-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "grpc_route", Usage: "The grpc_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("grpc_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetGrpcRouteRequest{Name: name}
									resp, err := client.GetGrpcRoute(ctx, req)
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
								Usage: "create grpc-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "grpc-route-id", Usage: "The grpc route id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateGrpcRouteRequest{Parent: parent}
									req.GrpcRouteId = cmd.String("grpc-route-id")
									req.GrpcRoute = &networkservicespb.GrpcRoute{
										Description: cmd.String("description"),
									}
									op, err := client.CreateGrpcRoute(ctx, req)
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
								Usage: "update grpc-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "grpc_route", Usage: "The grpc_route.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("grpc_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateGrpcRouteRequest{}
									req.GrpcRoute = &networkservicespb.GrpcRoute{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateGrpcRoute(ctx, req)
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
								Usage: "delete grpc-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "grpc_route", Usage: "The grpc_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/grpcRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("grpc_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteGrpcRouteRequest{Name: name}
									op, err := client.DeleteGrpcRoute(ctx, req)
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
						Name:  "http-routes",
						Usage: "Manage http-routes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list http-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListHttpRoutesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListHttpRoutes(ctx, req)
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
								Usage: "describe http-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "http_route", Usage: "The http_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/httpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("http_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetHttpRouteRequest{Name: name}
									resp, err := client.GetHttpRoute(ctx, req)
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
								Usage: "create http-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "http-route-id", Usage: "The http route id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateHttpRouteRequest{Parent: parent}
									req.HttpRouteId = cmd.String("http-route-id")
									req.HttpRoute = &networkservicespb.HttpRoute{
										Description: cmd.String("description"),
									}
									op, err := client.CreateHttpRoute(ctx, req)
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
								Usage: "update http-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "http_route", Usage: "The http_route.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/httpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("http_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateHttpRouteRequest{}
									req.HttpRoute = &networkservicespb.HttpRoute{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateHttpRoute(ctx, req)
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
								Usage: "delete http-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "http_route", Usage: "The http_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/httpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("http_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteHttpRouteRequest{Name: name}
									op, err := client.DeleteHttpRoute(ctx, req)
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
						Name:  "lb-edge-extensions",
						Usage: "Manage lb-edge-extensions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list lb-edge-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListLbEdgeExtensionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLbEdgeExtensions(ctx, req)
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
								Usage: "describe lb-edge-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_edge_extension", Usage: "The lb_edge_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbEdgeExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_edge_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetLbEdgeExtensionRequest{Name: name}
									resp, err := client.GetLbEdgeExtension(ctx, req)
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
								Usage: "create lb-edge-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb-edge-extension-id", Usage: "The lb edge extension id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateLbEdgeExtensionRequest{Parent: parent}
									req.LbEdgeExtensionId = cmd.String("lb-edge-extension-id")
									req.LbEdgeExtension = &networkservicespb.LbEdgeExtension{
										Description: cmd.String("description"),
									}
									op, err := client.CreateLbEdgeExtension(ctx, req)
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
								Usage: "update lb-edge-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_edge_extension", Usage: "The lb_edge_extension.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbEdgeExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_edge_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateLbEdgeExtensionRequest{}
									req.LbEdgeExtension = &networkservicespb.LbEdgeExtension{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateLbEdgeExtension(ctx, req)
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
								Usage: "delete lb-edge-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_edge_extension", Usage: "The lb_edge_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbEdgeExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_edge_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteLbEdgeExtensionRequest{Name: name}
									op, err := client.DeleteLbEdgeExtension(ctx, req)
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
						Name:  "lb-route-extensions",
						Usage: "Manage lb-route-extensions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list lb-route-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListLbRouteExtensionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLbRouteExtensions(ctx, req)
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
								Usage: "describe lb-route-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_route_extension", Usage: "The lb_route_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbRouteExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_route_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetLbRouteExtensionRequest{Name: name}
									resp, err := client.GetLbRouteExtension(ctx, req)
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
								Usage: "create lb-route-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb-route-extension-id", Usage: "The lb route extension id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateLbRouteExtensionRequest{Parent: parent}
									req.LbRouteExtensionId = cmd.String("lb-route-extension-id")
									req.LbRouteExtension = &networkservicespb.LbRouteExtension{
										Description: cmd.String("description"),
									}
									op, err := client.CreateLbRouteExtension(ctx, req)
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
								Usage: "update lb-route-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_route_extension", Usage: "The lb_route_extension.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbRouteExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_route_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateLbRouteExtensionRequest{}
									req.LbRouteExtension = &networkservicespb.LbRouteExtension{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateLbRouteExtension(ctx, req)
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
								Usage: "delete lb-route-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_route_extension", Usage: "The lb_route_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbRouteExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_route_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteLbRouteExtensionRequest{Name: name}
									op, err := client.DeleteLbRouteExtension(ctx, req)
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
						Name:  "lb-traffic-extensions",
						Usage: "Manage lb-traffic-extensions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list lb-traffic-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListLbTrafficExtensionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLbTrafficExtensions(ctx, req)
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
								Usage: "describe lb-traffic-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_traffic_extension", Usage: "The lb_traffic_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbTrafficExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_traffic_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetLbTrafficExtensionRequest{Name: name}
									resp, err := client.GetLbTrafficExtension(ctx, req)
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
								Usage: "create lb-traffic-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb-traffic-extension-id", Usage: "The lb traffic extension id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateLbTrafficExtensionRequest{Parent: parent}
									req.LbTrafficExtensionId = cmd.String("lb-traffic-extension-id")
									req.LbTrafficExtension = &networkservicespb.LbTrafficExtension{
										Description: cmd.String("description"),
									}
									op, err := client.CreateLbTrafficExtension(ctx, req)
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
								Usage: "update lb-traffic-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_traffic_extension", Usage: "The lb_traffic_extension.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbTrafficExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_traffic_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateLbTrafficExtensionRequest{}
									req.LbTrafficExtension = &networkservicespb.LbTrafficExtension{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateLbTrafficExtension(ctx, req)
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
								Usage: "delete lb-traffic-extensions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "lb_traffic_extension", Usage: "The lb_traffic_extension.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/lbTrafficExtensions/%s", cmd.String("project"), cmd.String("location"), cmd.String("lb_traffic_extension"))
									client, err := networkservices.NewDepClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteLbTrafficExtensionRequest{Name: name}
									op, err := client.DeleteLbTrafficExtension(ctx, req)
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
									client, err := networkservices.NewDepClient(ctx)
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
									client, err := networkservices.NewDepClient(ctx)
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
									client, err := networkservices.NewClient(ctx)
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
									client, err := networkservices.NewClient(ctx)
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
						Name:  "meshes",
						Usage: "Manage meshes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list meshes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListMeshesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMeshes(ctx, req)
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
								Usage: "describe meshes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mesh", Usage: "The mesh.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("mesh"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetMeshRequest{Name: name}
									resp, err := client.GetMesh(ctx, req)
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
								Usage: "create meshes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mesh-id", Usage: "The mesh id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.IntFlag{Name: "interception-port", Usage: "The interception port.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateMeshRequest{Parent: parent}
									req.MeshId = cmd.String("mesh-id")
									req.Mesh = &networkservicespb.Mesh{
										Description:      cmd.String("description"),
										InterceptionPort: int32(cmd.Int("interception-port")),
									}
									op, err := client.CreateMesh(ctx, req)
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
								Usage: "update meshes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mesh", Usage: "The mesh.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.IntFlag{Name: "interception-port", Usage: "The interception port.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("mesh"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateMeshRequest{}
									req.Mesh = &networkservicespb.Mesh{
										Name:             name,
										Description:      cmd.String("description"),
										InterceptionPort: int32(cmd.Int("interception-port")),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("interception-port") {
										paths = append(paths, "interception_port")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMesh(ctx, req)
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
								Usage: "delete meshes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mesh", Usage: "The mesh.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("mesh"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteMeshRequest{Name: name}
									op, err := client.DeleteMesh(ctx, req)
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
									client, err := networkservices.NewDepClient(ctx)
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
									client, err := networkservices.NewDepClient(ctx)
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
									client, err := networkservices.NewDepClient(ctx)
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
									client, err := networkservices.NewDepClient(ctx)
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
									client, err := networkservices.NewClient(ctx)
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
									client, err := networkservices.NewClient(ctx)
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
									client, err := networkservices.NewClient(ctx)
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
									client, err := networkservices.NewClient(ctx)
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
						Name:  "route-views",
						Usage: "Manage route-views resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe route-views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
									&cli.StringFlag{Name: "route_view", Usage: "The route_view.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gateways/%s/routeViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"), cmd.String("route_view"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetGatewayRouteViewRequest{Name: name}
									resp, err := client.GetGatewayRouteView(ctx, req)
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
								Usage: "describe route-views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mesh", Usage: "The mesh.", Required: true},
									&cli.StringFlag{Name: "route_view", Usage: "The route_view.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/meshes/%s/routeViews/%s", cmd.String("project"), cmd.String("location"), cmd.String("mesh"), cmd.String("route_view"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetMeshRouteViewRequest{Name: name}
									resp, err := client.GetMeshRouteView(ctx, req)
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
								Usage: "list route-views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway", Usage: "The gateway.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/gateways/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListGatewayRouteViewsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListGatewayRouteViews(ctx, req)
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
								Name:  "list",
								Usage: "list route-views",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mesh", Usage: "The mesh.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/meshes/%s", cmd.String("project"), cmd.String("location"), cmd.String("mesh"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListMeshRouteViewsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMeshRouteViews(ctx, req)
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
						Name:  "service-bindings",
						Usage: "Manage service-bindings resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list service-bindings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListServiceBindingsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListServiceBindings(ctx, req)
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
								Usage: "describe service-bindings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service_binding", Usage: "The service_binding.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serviceBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_binding"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetServiceBindingRequest{Name: name}
									resp, err := client.GetServiceBinding(ctx, req)
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
								Usage: "create service-bindings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service-binding-id", Usage: "The service binding id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateServiceBindingRequest{Parent: parent}
									req.ServiceBindingId = cmd.String("service-binding-id")
									req.ServiceBinding = &networkservicespb.ServiceBinding{
										Description: cmd.String("description"),
										Service:     cmd.String("service"),
									}
									op, err := client.CreateServiceBinding(ctx, req)
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
								Usage: "update service-bindings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service_binding", Usage: "The service_binding.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "service", Usage: "The service.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serviceBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_binding"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateServiceBindingRequest{}
									req.ServiceBinding = &networkservicespb.ServiceBinding{
										Name:        name,
										Description: cmd.String("description"),
										Service:     cmd.String("service"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("service") {
										paths = append(paths, "service")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateServiceBinding(ctx, req)
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
								Usage: "delete service-bindings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service_binding", Usage: "The service_binding.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serviceBindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_binding"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteServiceBindingRequest{Name: name}
									op, err := client.DeleteServiceBinding(ctx, req)
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
						Name:  "service-lb-policies",
						Usage: "Manage service-lb-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list service-lb-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListServiceLbPoliciesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListServiceLbPolicies(ctx, req)
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
								Usage: "describe service-lb-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service_lb_policy", Usage: "The service_lb_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serviceLbPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_lb_policy"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetServiceLbPolicyRequest{Name: name}
									resp, err := client.GetServiceLbPolicy(ctx, req)
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
								Usage: "create service-lb-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service-lb-policy-id", Usage: "The service lb policy id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateServiceLbPolicyRequest{Parent: parent}
									req.ServiceLbPolicyId = cmd.String("service-lb-policy-id")
									req.ServiceLbPolicy = &networkservicespb.ServiceLbPolicy{
										Description: cmd.String("description"),
									}
									op, err := client.CreateServiceLbPolicy(ctx, req)
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
								Usage: "update service-lb-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service_lb_policy", Usage: "The service_lb_policy.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serviceLbPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_lb_policy"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateServiceLbPolicyRequest{}
									req.ServiceLbPolicy = &networkservicespb.ServiceLbPolicy{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateServiceLbPolicy(ctx, req)
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
								Usage: "delete service-lb-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "service_lb_policy", Usage: "The service_lb_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serviceLbPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("service_lb_policy"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteServiceLbPolicyRequest{Name: name}
									op, err := client.DeleteServiceLbPolicy(ctx, req)
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
						Name:  "tcp-routes",
						Usage: "Manage tcp-routes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list tcp-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListTcpRoutesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTcpRoutes(ctx, req)
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
								Usage: "describe tcp-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tcp_route", Usage: "The tcp_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tcpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tcp_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetTcpRouteRequest{Name: name}
									resp, err := client.GetTcpRoute(ctx, req)
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
								Usage: "create tcp-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tcp-route-id", Usage: "The tcp route id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateTcpRouteRequest{Parent: parent}
									req.TcpRouteId = cmd.String("tcp-route-id")
									req.TcpRoute = &networkservicespb.TcpRoute{
										Description: cmd.String("description"),
									}
									op, err := client.CreateTcpRoute(ctx, req)
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
								Usage: "update tcp-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tcp_route", Usage: "The tcp_route.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tcpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tcp_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateTcpRouteRequest{}
									req.TcpRoute = &networkservicespb.TcpRoute{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateTcpRoute(ctx, req)
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
								Usage: "delete tcp-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tcp_route", Usage: "The tcp_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tcpRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tcp_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteTcpRouteRequest{Name: name}
									op, err := client.DeleteTcpRoute(ctx, req)
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
						Name:  "tls-routes",
						Usage: "Manage tls-routes resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list tls-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListTlsRoutesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTlsRoutes(ctx, req)
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
								Usage: "describe tls-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls_route", Usage: "The tls_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetTlsRouteRequest{Name: name}
									resp, err := client.GetTlsRoute(ctx, req)
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
								Usage: "create tls-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls-route-id", Usage: "The tls route id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateTlsRouteRequest{Parent: parent}
									req.TlsRouteId = cmd.String("tls-route-id")
									req.TlsRoute = &networkservicespb.TlsRoute{
										Description: cmd.String("description"),
									}
									op, err := client.CreateTlsRoute(ctx, req)
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
								Usage: "update tls-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls_route", Usage: "The tls_route.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateTlsRouteRequest{}
									req.TlsRoute = &networkservicespb.TlsRoute{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateTlsRoute(ctx, req)
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
								Usage: "delete tls-routes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls_route", Usage: "The tls_route.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tlsRoutes/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls_route"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteTlsRouteRequest{Name: name}
									op, err := client.DeleteTlsRoute(ctx, req)
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
						Name:  "versions",
						Usage: "Manage versions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin", Usage: "The wasm_plugin.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm_plugin"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListWasmPluginVersionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWasmPluginVersions(ctx, req)
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
								Usage: "describe versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin", Usage: "The wasm_plugin.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin_version", Usage: "The wasm_plugin_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm_plugin"), cmd.String("wasm_plugin_version"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetWasmPluginVersionRequest{Name: name}
									resp, err := client.GetWasmPluginVersion(ctx, req)
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
								Usage: "create versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin", Usage: "The wasm_plugin.", Required: true},
									&cli.StringFlag{Name: "wasm-plugin-version-id", Usage: "The wasm plugin version id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "image-uri", Usage: "The image uri.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm_plugin"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateWasmPluginVersionRequest{Parent: parent}
									req.WasmPluginVersionId = cmd.String("wasm-plugin-version-id")
									req.WasmPluginVersion = &networkservicespb.WasmPluginVersion{
										Description: cmd.String("description"),
										ImageUri:    cmd.String("image-uri"),
									}
									op, err := client.CreateWasmPluginVersion(ctx, req)
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
								Usage: "delete versions",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin", Usage: "The wasm_plugin.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin_version", Usage: "The wasm_plugin_version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s/versions/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm_plugin"), cmd.String("wasm_plugin_version"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteWasmPluginVersionRequest{Name: name}
									op, err := client.DeleteWasmPluginVersion(ctx, req)
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
						Name:  "wasm-plugins",
						Usage: "Manage wasm-plugins resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list wasm-plugins",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networkservicespb.ListWasmPluginsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListWasmPlugins(ctx, req)
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
								Usage: "describe wasm-plugins",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin", Usage: "The wasm_plugin.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm_plugin"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.GetWasmPluginRequest{Name: name}
									resp, err := client.GetWasmPlugin(ctx, req)
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
								Usage: "create wasm-plugins",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm-plugin-id", Usage: "The wasm plugin id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "main-version-id", Usage: "The main version id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.CreateWasmPluginRequest{Parent: parent}
									req.WasmPluginId = cmd.String("wasm-plugin-id")
									req.WasmPlugin = &networkservicespb.WasmPlugin{
										Description:   cmd.String("description"),
										MainVersionId: cmd.String("main-version-id"),
									}
									op, err := client.CreateWasmPlugin(ctx, req)
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
								Usage: "update wasm-plugins",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin", Usage: "The wasm_plugin.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "main-version-id", Usage: "The main version id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm_plugin"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.UpdateWasmPluginRequest{}
									req.WasmPlugin = &networkservicespb.WasmPlugin{
										Name:          name,
										Description:   cmd.String("description"),
										MainVersionId: cmd.String("main-version-id"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("main-version-id") {
										paths = append(paths, "main_version_id")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateWasmPlugin(ctx, req)
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
								Usage: "delete wasm-plugins",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "wasm_plugin", Usage: "The wasm_plugin.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/wasmPlugins/%s", cmd.String("project"), cmd.String("location"), cmd.String("wasm_plugin"))
									client, err := networkservices.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networkservicespb.DeleteWasmPluginRequest{Name: name}
									op, err := client.DeleteWasmPlugin(ctx, req)
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
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
