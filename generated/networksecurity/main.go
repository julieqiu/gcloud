package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	networksecurity "cloud.google.com/go/networksecurity/apiv1"
	"cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
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
				Name:  "networksecurity",
				Usage: "manage Network Security API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "address-groups",
						Usage: "Manage address-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
									client, err := networksecurity.NewAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetAddressGroupRequest{Name: name}
									resp, err := client.GetAddressGroup(ctx, req)
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
								Usage: "create address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address-group-id", Usage: "The address group id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.IntFlag{Name: "capacity", Usage: "The capacity.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateAddressGroupRequest{Parent: parent}
									req.AddressGroupId = cmd.String("address-group-id")
									req.AddressGroup = &networksecuritypb.AddressGroup{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Capacity:    int32(cmd.Int("capacity")),
									}
									op, err := client.CreateAddressGroup(ctx, req)
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
								Usage: "update address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.IntFlag{Name: "capacity", Usage: "The capacity.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
									client, err := networksecurity.NewAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateAddressGroupRequest{}
									req.AddressGroup = &networksecuritypb.AddressGroup{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Capacity:    int32(cmd.Int("capacity")),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("capacity") {
										paths = append(paths, "capacity")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAddressGroup(ctx, req)
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
								Name:  "add-items",
								Usage: "add-items address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-items...")
									return nil
								},
							},
							{
								Name:  "remove-items",
								Usage: "remove-items address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-items...")
									return nil
								},
							},
							{
								Name:  "clone-items",
								Usage: "clone-items address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing clone-items...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
									client, err := networksecurity.NewAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteAddressGroupRequest{Name: name}
									op, err := client.DeleteAddressGroup(ctx, req)
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
								Name:  "list",
								Usage: "list address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "list",
								Usage: "list address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetAddressGroupRequest{Name: name}
									resp, err := client.GetAddressGroup(ctx, req)
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
								Usage: "create address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address-group-id", Usage: "The address group id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.IntFlag{Name: "capacity", Usage: "The capacity.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateAddressGroupRequest{Parent: parent}
									req.AddressGroupId = cmd.String("address-group-id")
									req.AddressGroup = &networksecuritypb.AddressGroup{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Capacity:    int32(cmd.Int("capacity")),
									}
									op, err := client.CreateAddressGroup(ctx, req)
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
								Usage: "update address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.IntFlag{Name: "capacity", Usage: "The capacity.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateAddressGroupRequest{}
									req.AddressGroup = &networksecuritypb.AddressGroup{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Capacity:    int32(cmd.Int("capacity")),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("capacity") {
										paths = append(paths, "capacity")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAddressGroup(ctx, req)
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
								Name:  "add-items",
								Usage: "add-items address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing add-items...")
									return nil
								},
							},
							{
								Name:  "remove-items",
								Usage: "remove-items address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing remove-items...")
									return nil
								},
							},
							{
								Name:  "clone-items",
								Usage: "clone-items address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing clone-items...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete address-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteAddressGroupRequest{Name: name}
									op, err := client.DeleteAddressGroup(ctx, req)
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
								Name:  "list",
								Usage: "list address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions address-groups",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "authorization-policies",
						Usage: "Manage authorization-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list authorization-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe authorization-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authorization_policy", Usage: "The authorization_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authorization_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetAuthorizationPolicyRequest{Name: name}
									resp, err := client.GetAuthorizationPolicy(ctx, req)
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
								Usage: "create authorization-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authorization-policy-id", Usage: "The authorization policy id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateAuthorizationPolicyRequest{Parent: parent}
									req.AuthorizationPolicyId = cmd.String("authorization-policy-id")
									req.AuthorizationPolicy = &networksecuritypb.AuthorizationPolicy{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateAuthorizationPolicy(ctx, req)
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
								Usage: "update authorization-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authorization_policy", Usage: "The authorization_policy.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authorization_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateAuthorizationPolicyRequest{}
									req.AuthorizationPolicy = &networksecuritypb.AuthorizationPolicy{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAuthorizationPolicy(ctx, req)
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
								Usage: "delete authorization-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authorization_policy", Usage: "The authorization_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authorization_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteAuthorizationPolicyRequest{Name: name}
									op, err := client.DeleteAuthorizationPolicy(ctx, req)
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
						Name:  "authz-policies",
						Usage: "Manage authz-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list authz-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListAuthzPoliciesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListAuthzPolicies(ctx, req)
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
								Usage: "describe authz-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz_policy", Usage: "The authz_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetAuthzPolicyRequest{Name: name}
									resp, err := client.GetAuthzPolicy(ctx, req)
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
								Usage: "create authz-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz-policy-id", Usage: "The authz policy id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateAuthzPolicyRequest{Parent: parent}
									req.AuthzPolicyId = cmd.String("authz-policy-id")
									req.AuthzPolicy = &networksecuritypb.AuthzPolicy{
										Description: cmd.String("description"),
									}
									op, err := client.CreateAuthzPolicy(ctx, req)
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
								Usage: "update authz-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz_policy", Usage: "The authz_policy.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateAuthzPolicyRequest{}
									req.AuthzPolicy = &networksecuritypb.AuthzPolicy{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateAuthzPolicy(ctx, req)
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
								Usage: "delete authz-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "authz_policy", Usage: "The authz_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteAuthzPolicyRequest{Name: name}
									op, err := client.DeleteAuthzPolicy(ctx, req)
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
						Name:  "backend-authentication-configs",
						Usage: "Manage backend-authentication-configs resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list backend-authentication-configs",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe backend-authentication-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backend_authentication_config", Usage: "The backend_authentication_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("backend_authentication_config"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetBackendAuthenticationConfigRequest{Name: name}
									resp, err := client.GetBackendAuthenticationConfig(ctx, req)
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
								Usage: "create backend-authentication-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backend-authentication-config-id", Usage: "The backend authentication config id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "client-certificate", Usage: "The client certificate.", Required: false},
									&cli.StringFlag{Name: "trust-config", Usage: "The trust config.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateBackendAuthenticationConfigRequest{Parent: parent}
									req.BackendAuthenticationConfigId = cmd.String("backend-authentication-config-id")
									req.BackendAuthenticationConfig = &networksecuritypb.BackendAuthenticationConfig{
										Name:              cmd.String("name"),
										Description:       cmd.String("description"),
										ClientCertificate: cmd.String("client-certificate"),
										TrustConfig:       cmd.String("trust-config"),
									}
									op, err := client.CreateBackendAuthenticationConfig(ctx, req)
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
								Usage: "update backend-authentication-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backend_authentication_config", Usage: "The backend_authentication_config.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "client-certificate", Usage: "The client certificate.", Required: false},
									&cli.StringFlag{Name: "trust-config", Usage: "The trust config.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("backend_authentication_config"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateBackendAuthenticationConfigRequest{}
									req.BackendAuthenticationConfig = &networksecuritypb.BackendAuthenticationConfig{
										Name:              name,
										Name:              cmd.String("name"),
										Description:       cmd.String("description"),
										ClientCertificate: cmd.String("client-certificate"),
										TrustConfig:       cmd.String("trust-config"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("client-certificate") {
										paths = append(paths, "client_certificate")
									}
									if cmd.IsSet("trust-config") {
										paths = append(paths, "trust_config")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateBackendAuthenticationConfig(ctx, req)
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
								Usage: "delete backend-authentication-configs",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "backend_authentication_config", Usage: "The backend_authentication_config.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("backend_authentication_config"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteBackendAuthenticationConfigRequest{Name: name}
									op, err := client.DeleteBackendAuthenticationConfig(ctx, req)
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
						Name:  "client-tls-policies",
						Usage: "Manage client-tls-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list client-tls-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe client-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "client_tls_policy", Usage: "The client_tls_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("client_tls_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetClientTlsPolicyRequest{Name: name}
									resp, err := client.GetClientTlsPolicy(ctx, req)
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
								Usage: "create client-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "client-tls-policy-id", Usage: "The client tls policy id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "sni", Usage: "The sni.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateClientTlsPolicyRequest{Parent: parent}
									req.ClientTlsPolicyId = cmd.String("client-tls-policy-id")
									req.ClientTlsPolicy = &networksecuritypb.ClientTlsPolicy{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Sni:         cmd.String("sni"),
									}
									op, err := client.CreateClientTlsPolicy(ctx, req)
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
								Usage: "update client-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "client_tls_policy", Usage: "The client_tls_policy.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "sni", Usage: "The sni.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("client_tls_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateClientTlsPolicyRequest{}
									req.ClientTlsPolicy = &networksecuritypb.ClientTlsPolicy{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										Sni:         cmd.String("sni"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("sni") {
										paths = append(paths, "sni")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateClientTlsPolicy(ctx, req)
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
								Usage: "delete client-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "client_tls_policy", Usage: "The client_tls_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("client_tls_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteClientTlsPolicyRequest{Name: name}
									op, err := client.DeleteClientTlsPolicy(ctx, req)
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
						Name:  "dns-threat-detectors",
						Usage: "Manage dns-threat-detectors resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list dns-threat-detectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListDnsThreatDetectorsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListDnsThreatDetectors(ctx, req)
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
								Usage: "describe dns-threat-detectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "dns_threat_detector", Usage: "The dns_threat_detector.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns_threat_detector"))
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetDnsThreatDetectorRequest{Name: name}
									resp, err := client.GetDnsThreatDetector(ctx, req)
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
								Usage: "create dns-threat-detectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "dns-threat-detector-id", Usage: "The dns threat detector id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateDnsThreatDetectorRequest{Parent: parent}
									req.DnsThreatDetectorId = cmd.String("dns-threat-detector-id")
									resp, err := client.CreateDnsThreatDetector(ctx, req)
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
								Usage: "update dns-threat-detectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "dns_threat_detector", Usage: "The dns_threat_detector.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns_threat_detector"))
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateDnsThreatDetectorRequest{}
									req.DnsThreatDetector = &networksecuritypb.DnsThreatDetector{
										Name: name,
									}
									resp, err := client.UpdateDnsThreatDetector(ctx, req)
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
								Usage: "delete dns-threat-detectors",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "dns_threat_detector", Usage: "The dns_threat_detector.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns_threat_detector"))
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteDnsThreatDetectorRequest{Name: name}
									if err := client.DeleteDnsThreatDetector(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
									return nil
								},
							},
						},
					},
					{
						Name:  "firewall-endpoint-associations",
						Usage: "Manage firewall-endpoint-associations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list firewall-endpoint-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListFirewallEndpointAssociationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListFirewallEndpointAssociations(ctx, req)
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
								Usage: "describe firewall-endpoint-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall_endpoint_association", Usage: "The firewall_endpoint_association.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall_endpoint_association"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetFirewallEndpointAssociationRequest{Name: name}
									resp, err := client.GetFirewallEndpointAssociation(ctx, req)
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
								Usage: "create firewall-endpoint-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall-endpoint-association-id", Usage: "The firewall endpoint association id.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
									&cli.StringFlag{Name: "firewall-endpoint", Usage: "The firewall endpoint.", Required: true},
									&cli.StringFlag{Name: "tls-inspection-policy", Usage: "The tls inspection policy.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateFirewallEndpointAssociationRequest{Parent: parent}
									req.FirewallEndpointAssociationId = cmd.String("firewall-endpoint-association-id")
									req.FirewallEndpointAssociation = &networksecuritypb.FirewallEndpointAssociation{
										Network:             cmd.String("network"),
										FirewallEndpoint:    cmd.String("firewall-endpoint"),
										TlsInspectionPolicy: cmd.String("tls-inspection-policy"),
										Disabled:            cmd.Bool("disabled"),
									}
									op, err := client.CreateFirewallEndpointAssociation(ctx, req)
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
								Usage: "delete firewall-endpoint-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall_endpoint_association", Usage: "The firewall_endpoint_association.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall_endpoint_association"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteFirewallEndpointAssociationRequest{Name: name}
									op, err := client.DeleteFirewallEndpointAssociation(ctx, req)
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
								Usage: "update firewall-endpoint-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall_endpoint_association", Usage: "The firewall_endpoint_association.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.StringFlag{Name: "firewall-endpoint", Usage: "The firewall endpoint.", Required: false},
									&cli.StringFlag{Name: "tls-inspection-policy", Usage: "The tls inspection policy.", Required: false},
									&cli.BoolFlag{Name: "disabled", Usage: "The disabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall_endpoint_association"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateFirewallEndpointAssociationRequest{}
									req.FirewallEndpointAssociation = &networksecuritypb.FirewallEndpointAssociation{
										Name:                name,
										Network:             cmd.String("network"),
										FirewallEndpoint:    cmd.String("firewall-endpoint"),
										TlsInspectionPolicy: cmd.String("tls-inspection-policy"),
										Disabled:            cmd.Bool("disabled"),
									}
									var paths []string
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									if cmd.IsSet("firewall-endpoint") {
										paths = append(paths, "firewall_endpoint")
									}
									if cmd.IsSet("tls-inspection-policy") {
										paths = append(paths, "tls_inspection_policy")
									}
									if cmd.IsSet("disabled") {
										paths = append(paths, "disabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateFirewallEndpointAssociation(ctx, req)
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
						Name:  "firewall-endpoints",
						Usage: "Manage firewall-endpoints resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list firewall-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListFirewallEndpointsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListFirewallEndpoints(ctx, req)
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
								Usage: "describe firewall-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall_endpoint", Usage: "The firewall_endpoint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall_endpoint"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetFirewallEndpointRequest{Name: name}
									resp, err := client.GetFirewallEndpoint(ctx, req)
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
								Usage: "create firewall-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall-endpoint-id", Usage: "The firewall endpoint id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "billing-project-id", Usage: "The billing project id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateFirewallEndpointRequest{Parent: parent}
									req.FirewallEndpointId = cmd.String("firewall-endpoint-id")
									req.FirewallEndpoint = &networksecuritypb.FirewallEndpoint{
										Description:      cmd.String("description"),
										BillingProjectId: cmd.String("billing-project-id"),
									}
									op, err := client.CreateFirewallEndpoint(ctx, req)
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
								Usage: "delete firewall-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall_endpoint", Usage: "The firewall_endpoint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall_endpoint"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteFirewallEndpointRequest{Name: name}
									op, err := client.DeleteFirewallEndpoint(ctx, req)
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
								Usage: "update firewall-endpoints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "firewall_endpoint", Usage: "The firewall_endpoint.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "billing-project-id", Usage: "The billing project id.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall_endpoint"))
									client, err := networksecurity.NewFirewallActivationClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateFirewallEndpointRequest{}
									req.FirewallEndpoint = &networksecuritypb.FirewallEndpoint{
										Name:             name,
										Description:      cmd.String("description"),
										BillingProjectId: cmd.String("billing-project-id"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("billing-project-id") {
										paths = append(paths, "billing_project_id")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateFirewallEndpoint(ctx, req)
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
						Name:  "gateway-security-policies",
						Usage: "Manage gateway-security-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list gateway-security-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe gateway-security-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetGatewaySecurityPolicyRequest{Name: name}
									resp, err := client.GetGatewaySecurityPolicy(ctx, req)
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
								Usage: "create gateway-security-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway-security-policy-id", Usage: "The gateway security policy id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "tls-inspection-policy", Usage: "The tls inspection policy.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateGatewaySecurityPolicyRequest{Parent: parent}
									req.GatewaySecurityPolicyId = cmd.String("gateway-security-policy-id")
									req.GatewaySecurityPolicy = &networksecuritypb.GatewaySecurityPolicy{
										Name:                cmd.String("name"),
										Description:         cmd.String("description"),
										TlsInspectionPolicy: cmd.String("tls-inspection-policy"),
									}
									op, err := client.CreateGatewaySecurityPolicy(ctx, req)
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
								Usage: "update gateway-security-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "tls-inspection-policy", Usage: "The tls inspection policy.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateGatewaySecurityPolicyRequest{}
									req.GatewaySecurityPolicy = &networksecuritypb.GatewaySecurityPolicy{
										Name:                name,
										Name:                cmd.String("name"),
										Description:         cmd.String("description"),
										TlsInspectionPolicy: cmd.String("tls-inspection-policy"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("tls-inspection-policy") {
										paths = append(paths, "tls_inspection_policy")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateGatewaySecurityPolicy(ctx, req)
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
								Usage: "delete gateway-security-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteGatewaySecurityPolicyRequest{Name: name}
									op, err := client.DeleteGatewaySecurityPolicy(ctx, req)
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
						Name:  "intercept-deployment-groups",
						Usage: "Manage intercept-deployment-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list intercept-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListInterceptDeploymentGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInterceptDeploymentGroups(ctx, req)
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
								Usage: "describe intercept-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_deployment_group", Usage: "The intercept_deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment_group"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetInterceptDeploymentGroupRequest{Name: name}
									resp, err := client.GetInterceptDeploymentGroup(ctx, req)
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
								Usage: "create intercept-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept-deployment-group-id", Usage: "The intercept deployment group id.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateInterceptDeploymentGroupRequest{Parent: parent}
									req.InterceptDeploymentGroupId = cmd.String("intercept-deployment-group-id")
									req.InterceptDeploymentGroup = &networksecuritypb.InterceptDeploymentGroup{
										Network:     cmd.String("network"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateInterceptDeploymentGroup(ctx, req)
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
								Usage: "update intercept-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_deployment_group", Usage: "The intercept_deployment_group.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment_group"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateInterceptDeploymentGroupRequest{}
									req.InterceptDeploymentGroup = &networksecuritypb.InterceptDeploymentGroup{
										Name:        name,
										Network:     cmd.String("network"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateInterceptDeploymentGroup(ctx, req)
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
								Usage: "delete intercept-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_deployment_group", Usage: "The intercept_deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment_group"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteInterceptDeploymentGroupRequest{Name: name}
									op, err := client.DeleteInterceptDeploymentGroup(ctx, req)
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
						Name:  "intercept-deployments",
						Usage: "Manage intercept-deployments resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list intercept-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListInterceptDeploymentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInterceptDeployments(ctx, req)
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
								Usage: "describe intercept-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_deployment", Usage: "The intercept_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetInterceptDeploymentRequest{Name: name}
									resp, err := client.GetInterceptDeployment(ctx, req)
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
								Usage: "create intercept-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept-deployment-id", Usage: "The intercept deployment id.", Required: true},
									&cli.StringFlag{Name: "forwarding-rule", Usage: "The forwarding rule.", Required: true},
									&cli.StringFlag{Name: "intercept-deployment-group", Usage: "The intercept deployment group.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateInterceptDeploymentRequest{Parent: parent}
									req.InterceptDeploymentId = cmd.String("intercept-deployment-id")
									req.InterceptDeployment = &networksecuritypb.InterceptDeployment{
										ForwardingRule:           cmd.String("forwarding-rule"),
										InterceptDeploymentGroup: cmd.String("intercept-deployment-group"),
										Description:              cmd.String("description"),
									}
									op, err := client.CreateInterceptDeployment(ctx, req)
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
								Usage: "update intercept-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_deployment", Usage: "The intercept_deployment.", Required: true},
									&cli.StringFlag{Name: "forwarding-rule", Usage: "The forwarding rule.", Required: false},
									&cli.StringFlag{Name: "intercept-deployment-group", Usage: "The intercept deployment group.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateInterceptDeploymentRequest{}
									req.InterceptDeployment = &networksecuritypb.InterceptDeployment{
										Name:                     name,
										ForwardingRule:           cmd.String("forwarding-rule"),
										InterceptDeploymentGroup: cmd.String("intercept-deployment-group"),
										Description:              cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("forwarding-rule") {
										paths = append(paths, "forwarding_rule")
									}
									if cmd.IsSet("intercept-deployment-group") {
										paths = append(paths, "intercept_deployment_group")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateInterceptDeployment(ctx, req)
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
								Usage: "delete intercept-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_deployment", Usage: "The intercept_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteInterceptDeploymentRequest{Name: name}
									op, err := client.DeleteInterceptDeployment(ctx, req)
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
						Name:  "intercept-endpoint-group-associations",
						Usage: "Manage intercept-endpoint-group-associations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list intercept-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListInterceptEndpointGroupAssociationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInterceptEndpointGroupAssociations(ctx, req)
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
								Usage: "describe intercept-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_endpoint_group_association", Usage: "The intercept_endpoint_group_association.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group_association"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetInterceptEndpointGroupAssociationRequest{Name: name}
									resp, err := client.GetInterceptEndpointGroupAssociation(ctx, req)
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
								Usage: "create intercept-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept-endpoint-group-association-id", Usage: "The intercept endpoint group association id.", Required: false},
									&cli.StringFlag{Name: "intercept-endpoint-group", Usage: "The intercept endpoint group.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateInterceptEndpointGroupAssociationRequest{Parent: parent}
									req.InterceptEndpointGroupAssociationId = cmd.String("intercept-endpoint-group-association-id")
									req.InterceptEndpointGroupAssociation = &networksecuritypb.InterceptEndpointGroupAssociation{
										InterceptEndpointGroup: cmd.String("intercept-endpoint-group"),
										Network:                cmd.String("network"),
									}
									op, err := client.CreateInterceptEndpointGroupAssociation(ctx, req)
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
								Usage: "update intercept-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_endpoint_group_association", Usage: "The intercept_endpoint_group_association.", Required: true},
									&cli.StringFlag{Name: "intercept-endpoint-group", Usage: "The intercept endpoint group.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group_association"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateInterceptEndpointGroupAssociationRequest{}
									req.InterceptEndpointGroupAssociation = &networksecuritypb.InterceptEndpointGroupAssociation{
										Name:                   name,
										InterceptEndpointGroup: cmd.String("intercept-endpoint-group"),
										Network:                cmd.String("network"),
									}
									var paths []string
									if cmd.IsSet("intercept-endpoint-group") {
										paths = append(paths, "intercept_endpoint_group")
									}
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateInterceptEndpointGroupAssociation(ctx, req)
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
								Usage: "delete intercept-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_endpoint_group_association", Usage: "The intercept_endpoint_group_association.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group_association"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteInterceptEndpointGroupAssociationRequest{Name: name}
									op, err := client.DeleteInterceptEndpointGroupAssociation(ctx, req)
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
						Name:  "intercept-endpoint-groups",
						Usage: "Manage intercept-endpoint-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list intercept-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListInterceptEndpointGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListInterceptEndpointGroups(ctx, req)
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
								Usage: "describe intercept-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_endpoint_group", Usage: "The intercept_endpoint_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetInterceptEndpointGroupRequest{Name: name}
									resp, err := client.GetInterceptEndpointGroup(ctx, req)
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
								Usage: "create intercept-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept-endpoint-group-id", Usage: "The intercept endpoint group id.", Required: true},
									&cli.StringFlag{Name: "intercept-deployment-group", Usage: "The intercept deployment group.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateInterceptEndpointGroupRequest{Parent: parent}
									req.InterceptEndpointGroupId = cmd.String("intercept-endpoint-group-id")
									req.InterceptEndpointGroup = &networksecuritypb.InterceptEndpointGroup{
										InterceptDeploymentGroup: cmd.String("intercept-deployment-group"),
										Description:              cmd.String("description"),
									}
									op, err := client.CreateInterceptEndpointGroup(ctx, req)
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
								Usage: "update intercept-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_endpoint_group", Usage: "The intercept_endpoint_group.", Required: true},
									&cli.StringFlag{Name: "intercept-deployment-group", Usage: "The intercept deployment group.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateInterceptEndpointGroupRequest{}
									req.InterceptEndpointGroup = &networksecuritypb.InterceptEndpointGroup{
										Name:                     name,
										InterceptDeploymentGroup: cmd.String("intercept-deployment-group"),
										Description:              cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("intercept-deployment-group") {
										paths = append(paths, "intercept_deployment_group")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateInterceptEndpointGroup(ctx, req)
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
								Usage: "delete intercept-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "intercept_endpoint_group", Usage: "The intercept_endpoint_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group"))
									client, err := networksecurity.NewInterceptClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteInterceptEndpointGroupRequest{Name: name}
									op, err := client.DeleteInterceptEndpointGroup(ctx, req)
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
									client, err := networksecurity.NewAddressGroupClient(ctx)
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
									client, err := networksecurity.NewAddressGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
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
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
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
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
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
									client, err := networksecurity.NewFirewallActivationClient(ctx)
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
									client, err := networksecurity.NewFirewallActivationClient(ctx)
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
									client, err := networksecurity.NewInterceptClient(ctx)
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
									client, err := networksecurity.NewInterceptClient(ctx)
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
									client, err := networksecurity.NewMirroringClient(ctx)
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
									client, err := networksecurity.NewMirroringClient(ctx)
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
									client, err := networksecurity.NewClient(ctx)
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
									client, err := networksecurity.NewClient(ctx)
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
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
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
						Name:  "mirroring-deployment-groups",
						Usage: "Manage mirroring-deployment-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list mirroring-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListMirroringDeploymentGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMirroringDeploymentGroups(ctx, req)
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
								Usage: "describe mirroring-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_deployment_group", Usage: "The mirroring_deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment_group"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetMirroringDeploymentGroupRequest{Name: name}
									resp, err := client.GetMirroringDeploymentGroup(ctx, req)
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
								Usage: "create mirroring-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring-deployment-group-id", Usage: "The mirroring deployment group id.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateMirroringDeploymentGroupRequest{Parent: parent}
									req.MirroringDeploymentGroupId = cmd.String("mirroring-deployment-group-id")
									req.MirroringDeploymentGroup = &networksecuritypb.MirroringDeploymentGroup{
										Network:     cmd.String("network"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateMirroringDeploymentGroup(ctx, req)
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
								Usage: "update mirroring-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_deployment_group", Usage: "The mirroring_deployment_group.", Required: true},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment_group"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateMirroringDeploymentGroupRequest{}
									req.MirroringDeploymentGroup = &networksecuritypb.MirroringDeploymentGroup{
										Name:        name,
										Network:     cmd.String("network"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMirroringDeploymentGroup(ctx, req)
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
								Usage: "delete mirroring-deployment-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_deployment_group", Usage: "The mirroring_deployment_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment_group"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteMirroringDeploymentGroupRequest{Name: name}
									op, err := client.DeleteMirroringDeploymentGroup(ctx, req)
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
						Name:  "mirroring-deployments",
						Usage: "Manage mirroring-deployments resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list mirroring-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListMirroringDeploymentsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMirroringDeployments(ctx, req)
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
								Usage: "describe mirroring-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_deployment", Usage: "The mirroring_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetMirroringDeploymentRequest{Name: name}
									resp, err := client.GetMirroringDeployment(ctx, req)
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
								Usage: "create mirroring-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring-deployment-id", Usage: "The mirroring deployment id.", Required: true},
									&cli.StringFlag{Name: "forwarding-rule", Usage: "The forwarding rule.", Required: true},
									&cli.StringFlag{Name: "mirroring-deployment-group", Usage: "The mirroring deployment group.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateMirroringDeploymentRequest{Parent: parent}
									req.MirroringDeploymentId = cmd.String("mirroring-deployment-id")
									req.MirroringDeployment = &networksecuritypb.MirroringDeployment{
										ForwardingRule:           cmd.String("forwarding-rule"),
										MirroringDeploymentGroup: cmd.String("mirroring-deployment-group"),
										Description:              cmd.String("description"),
									}
									op, err := client.CreateMirroringDeployment(ctx, req)
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
								Usage: "update mirroring-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_deployment", Usage: "The mirroring_deployment.", Required: true},
									&cli.StringFlag{Name: "forwarding-rule", Usage: "The forwarding rule.", Required: false},
									&cli.StringFlag{Name: "mirroring-deployment-group", Usage: "The mirroring deployment group.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateMirroringDeploymentRequest{}
									req.MirroringDeployment = &networksecuritypb.MirroringDeployment{
										Name:                     name,
										ForwardingRule:           cmd.String("forwarding-rule"),
										MirroringDeploymentGroup: cmd.String("mirroring-deployment-group"),
										Description:              cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("forwarding-rule") {
										paths = append(paths, "forwarding_rule")
									}
									if cmd.IsSet("mirroring-deployment-group") {
										paths = append(paths, "mirroring_deployment_group")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMirroringDeployment(ctx, req)
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
								Usage: "delete mirroring-deployments",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_deployment", Usage: "The mirroring_deployment.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteMirroringDeploymentRequest{Name: name}
									op, err := client.DeleteMirroringDeployment(ctx, req)
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
						Name:  "mirroring-endpoint-group-associations",
						Usage: "Manage mirroring-endpoint-group-associations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list mirroring-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListMirroringEndpointGroupAssociationsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMirroringEndpointGroupAssociations(ctx, req)
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
								Usage: "describe mirroring-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_endpoint_group_association", Usage: "The mirroring_endpoint_group_association.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group_association"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetMirroringEndpointGroupAssociationRequest{Name: name}
									resp, err := client.GetMirroringEndpointGroupAssociation(ctx, req)
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
								Usage: "create mirroring-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring-endpoint-group-association-id", Usage: "The mirroring endpoint group association id.", Required: false},
									&cli.StringFlag{Name: "mirroring-endpoint-group", Usage: "The mirroring endpoint group.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateMirroringEndpointGroupAssociationRequest{Parent: parent}
									req.MirroringEndpointGroupAssociationId = cmd.String("mirroring-endpoint-group-association-id")
									req.MirroringEndpointGroupAssociation = &networksecuritypb.MirroringEndpointGroupAssociation{
										MirroringEndpointGroup: cmd.String("mirroring-endpoint-group"),
										Network:                cmd.String("network"),
									}
									op, err := client.CreateMirroringEndpointGroupAssociation(ctx, req)
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
								Usage: "update mirroring-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_endpoint_group_association", Usage: "The mirroring_endpoint_group_association.", Required: true},
									&cli.StringFlag{Name: "mirroring-endpoint-group", Usage: "The mirroring endpoint group.", Required: false},
									&cli.StringFlag{Name: "network", Usage: "The network.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group_association"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateMirroringEndpointGroupAssociationRequest{}
									req.MirroringEndpointGroupAssociation = &networksecuritypb.MirroringEndpointGroupAssociation{
										Name:                   name,
										MirroringEndpointGroup: cmd.String("mirroring-endpoint-group"),
										Network:                cmd.String("network"),
									}
									var paths []string
									if cmd.IsSet("mirroring-endpoint-group") {
										paths = append(paths, "mirroring_endpoint_group")
									}
									if cmd.IsSet("network") {
										paths = append(paths, "network")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMirroringEndpointGroupAssociation(ctx, req)
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
								Usage: "delete mirroring-endpoint-group-associations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_endpoint_group_association", Usage: "The mirroring_endpoint_group_association.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group_association"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteMirroringEndpointGroupAssociationRequest{Name: name}
									op, err := client.DeleteMirroringEndpointGroupAssociation(ctx, req)
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
						Name:  "mirroring-endpoint-groups",
						Usage: "Manage mirroring-endpoint-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list mirroring-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListMirroringEndpointGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMirroringEndpointGroups(ctx, req)
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
								Usage: "describe mirroring-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_endpoint_group", Usage: "The mirroring_endpoint_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetMirroringEndpointGroupRequest{Name: name}
									resp, err := client.GetMirroringEndpointGroup(ctx, req)
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
								Usage: "create mirroring-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring-endpoint-group-id", Usage: "The mirroring endpoint group id.", Required: true},
									&cli.StringFlag{Name: "mirroring-deployment-group", Usage: "The mirroring deployment group.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateMirroringEndpointGroupRequest{Parent: parent}
									req.MirroringEndpointGroupId = cmd.String("mirroring-endpoint-group-id")
									req.MirroringEndpointGroup = &networksecuritypb.MirroringEndpointGroup{
										MirroringDeploymentGroup: cmd.String("mirroring-deployment-group"),
										Description:              cmd.String("description"),
									}
									op, err := client.CreateMirroringEndpointGroup(ctx, req)
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
								Usage: "update mirroring-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_endpoint_group", Usage: "The mirroring_endpoint_group.", Required: true},
									&cli.StringFlag{Name: "mirroring-deployment-group", Usage: "The mirroring deployment group.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateMirroringEndpointGroupRequest{}
									req.MirroringEndpointGroup = &networksecuritypb.MirroringEndpointGroup{
										Name:                     name,
										MirroringDeploymentGroup: cmd.String("mirroring-deployment-group"),
										Description:              cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("mirroring-deployment-group") {
										paths = append(paths, "mirroring_deployment_group")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateMirroringEndpointGroup(ctx, req)
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
								Usage: "delete mirroring-endpoint-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "mirroring_endpoint_group", Usage: "The mirroring_endpoint_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group"))
									client, err := networksecurity.NewMirroringClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteMirroringEndpointGroupRequest{Name: name}
									op, err := client.DeleteMirroringEndpointGroup(ctx, req)
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
									client, err := networksecurity.NewAddressGroupClient(ctx)
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
									client, err := networksecurity.NewAddressGroupClient(ctx)
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
									client, err := networksecurity.NewAddressGroupClient(ctx)
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
									client, err := networksecurity.NewAddressGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationAddressGroupClient(ctx)
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
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
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
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
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
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
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
									client, err := networksecurity.NewDnsThreatDetectorClient(ctx)
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
									client, err := networksecurity.NewFirewallActivationClient(ctx)
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
									client, err := networksecurity.NewFirewallActivationClient(ctx)
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
									client, err := networksecurity.NewFirewallActivationClient(ctx)
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
									client, err := networksecurity.NewFirewallActivationClient(ctx)
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
									client, err := networksecurity.NewInterceptClient(ctx)
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
									client, err := networksecurity.NewInterceptClient(ctx)
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
									client, err := networksecurity.NewInterceptClient(ctx)
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
									client, err := networksecurity.NewInterceptClient(ctx)
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
									client, err := networksecurity.NewMirroringClient(ctx)
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
									client, err := networksecurity.NewMirroringClient(ctx)
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
									client, err := networksecurity.NewMirroringClient(ctx)
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
									client, err := networksecurity.NewMirroringClient(ctx)
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
									client, err := networksecurity.NewClient(ctx)
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
									client, err := networksecurity.NewClient(ctx)
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
									client, err := networksecurity.NewClient(ctx)
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
									client, err := networksecurity.NewClient(ctx)
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
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
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
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
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
						Name:  "rules",
						Usage: "Manage rules resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListGatewaySecurityPolicyRulesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListGatewaySecurityPolicyRules(ctx, req)
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
								Usage: "describe rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
									&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"), cmd.String("rule"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetGatewaySecurityPolicyRuleRequest{Name: name}
									resp, err := client.GetGatewaySecurityPolicyRule(ctx, req)
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
								Usage: "create rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
									&cli.StringFlag{Name: "gateway-security-policy-rule-id", Usage: "The gateway security policy rule id.", Required: false},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.BoolFlag{Name: "enabled", Usage: "The enabled.", Required: true},
									&cli.IntFlag{Name: "priority", Usage: "The priority.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "session-matcher", Usage: "The session matcher.", Required: true},
									&cli.StringFlag{Name: "application-matcher", Usage: "The application matcher.", Required: false},
									&cli.BoolFlag{Name: "tls-inspection-enabled", Usage: "The tls inspection enabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateGatewaySecurityPolicyRuleRequest{Parent: parent}
									req.GatewaySecurityPolicyRuleId = cmd.String("gateway-security-policy-rule-id")
									req.GatewaySecurityPolicyRule = &networksecuritypb.GatewaySecurityPolicyRule{
										Name:                 cmd.String("name"),
										Enabled:              cmd.Bool("enabled"),
										Priority:             int32(cmd.Int("priority")),
										Description:          cmd.String("description"),
										SessionMatcher:       cmd.String("session-matcher"),
										ApplicationMatcher:   cmd.String("application-matcher"),
										TlsInspectionEnabled: cmd.Bool("tls-inspection-enabled"),
									}
									op, err := client.CreateGatewaySecurityPolicyRule(ctx, req)
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
								Usage: "update rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
									&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.BoolFlag{Name: "enabled", Usage: "The enabled.", Required: false},
									&cli.IntFlag{Name: "priority", Usage: "The priority.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "session-matcher", Usage: "The session matcher.", Required: false},
									&cli.StringFlag{Name: "application-matcher", Usage: "The application matcher.", Required: false},
									&cli.BoolFlag{Name: "tls-inspection-enabled", Usage: "The tls inspection enabled.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"), cmd.String("rule"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateGatewaySecurityPolicyRuleRequest{}
									req.GatewaySecurityPolicyRule = &networksecuritypb.GatewaySecurityPolicyRule{
										Name:                 name,
										Name:                 cmd.String("name"),
										Enabled:              cmd.Bool("enabled"),
										Priority:             int32(cmd.Int("priority")),
										Description:          cmd.String("description"),
										SessionMatcher:       cmd.String("session-matcher"),
										ApplicationMatcher:   cmd.String("application-matcher"),
										TlsInspectionEnabled: cmd.Bool("tls-inspection-enabled"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("enabled") {
										paths = append(paths, "enabled")
									}
									if cmd.IsSet("priority") {
										paths = append(paths, "priority")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("session-matcher") {
										paths = append(paths, "session_matcher")
									}
									if cmd.IsSet("application-matcher") {
										paths = append(paths, "application_matcher")
									}
									if cmd.IsSet("tls-inspection-enabled") {
										paths = append(paths, "tls_inspection_enabled")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateGatewaySecurityPolicyRule(ctx, req)
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
								Usage: "delete rules",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
									&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"), cmd.String("rule"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteGatewaySecurityPolicyRuleRequest{Name: name}
									op, err := client.DeleteGatewaySecurityPolicyRule(ctx, req)
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
						Name:  "security-profile-groups",
						Usage: "Manage security-profile-groups resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list security-profile-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListSecurityProfileGroupsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSecurityProfileGroups(ctx, req)
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
								Usage: "describe security-profile-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_profile_group", Usage: "The security_profile_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/securityProfileGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile_group"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetSecurityProfileGroupRequest{Name: name}
									resp, err := client.GetSecurityProfileGroup(ctx, req)
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
								Usage: "create security-profile-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security-profile-group-id", Usage: "The security profile group id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "threat-prevention-profile", Usage: "The threat prevention profile.", Required: false},
									&cli.StringFlag{Name: "custom-mirroring-profile", Usage: "The custom mirroring profile.", Required: false},
									&cli.StringFlag{Name: "custom-intercept-profile", Usage: "The custom intercept profile.", Required: false},
									&cli.StringFlag{Name: "url-filtering-profile", Usage: "The url filtering profile.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateSecurityProfileGroupRequest{Parent: parent}
									req.SecurityProfileGroupId = cmd.String("security-profile-group-id")
									req.SecurityProfileGroup = &networksecuritypb.SecurityProfileGroup{
										Description:             cmd.String("description"),
										ThreatPreventionProfile: cmd.String("threat-prevention-profile"),
										CustomMirroringProfile:  cmd.String("custom-mirroring-profile"),
										CustomInterceptProfile:  cmd.String("custom-intercept-profile"),
										UrlFilteringProfile:     cmd.String("url-filtering-profile"),
									}
									op, err := client.CreateSecurityProfileGroup(ctx, req)
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
								Usage: "update security-profile-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_profile_group", Usage: "The security_profile_group.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "threat-prevention-profile", Usage: "The threat prevention profile.", Required: false},
									&cli.StringFlag{Name: "custom-mirroring-profile", Usage: "The custom mirroring profile.", Required: false},
									&cli.StringFlag{Name: "custom-intercept-profile", Usage: "The custom intercept profile.", Required: false},
									&cli.StringFlag{Name: "url-filtering-profile", Usage: "The url filtering profile.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/securityProfileGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile_group"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateSecurityProfileGroupRequest{}
									req.SecurityProfileGroup = &networksecuritypb.SecurityProfileGroup{
										Name:                    name,
										Description:             cmd.String("description"),
										ThreatPreventionProfile: cmd.String("threat-prevention-profile"),
										CustomMirroringProfile:  cmd.String("custom-mirroring-profile"),
										CustomInterceptProfile:  cmd.String("custom-intercept-profile"),
										UrlFilteringProfile:     cmd.String("url-filtering-profile"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("threat-prevention-profile") {
										paths = append(paths, "threat_prevention_profile")
									}
									if cmd.IsSet("custom-mirroring-profile") {
										paths = append(paths, "custom_mirroring_profile")
									}
									if cmd.IsSet("custom-intercept-profile") {
										paths = append(paths, "custom_intercept_profile")
									}
									if cmd.IsSet("url-filtering-profile") {
										paths = append(paths, "url_filtering_profile")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateSecurityProfileGroup(ctx, req)
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
								Usage: "delete security-profile-groups",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_profile_group", Usage: "The security_profile_group.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/securityProfileGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile_group"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteSecurityProfileGroupRequest{Name: name}
									op, err := client.DeleteSecurityProfileGroup(ctx, req)
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
						Name:  "security-profiles",
						Usage: "Manage security-profiles resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list security-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &networksecuritypb.ListSecurityProfilesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListSecurityProfiles(ctx, req)
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
								Usage: "describe security-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_profile", Usage: "The security_profile.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/securityProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetSecurityProfileRequest{Name: name}
									resp, err := client.GetSecurityProfile(ctx, req)
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
								Usage: "create security-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security-profile-id", Usage: "The security profile id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateSecurityProfileRequest{Parent: parent}
									req.SecurityProfileId = cmd.String("security-profile-id")
									req.SecurityProfile = &networksecuritypb.SecurityProfile{
										Description: cmd.String("description"),
									}
									op, err := client.CreateSecurityProfile(ctx, req)
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
								Usage: "update security-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_profile", Usage: "The security_profile.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/securityProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateSecurityProfileRequest{}
									req.SecurityProfile = &networksecuritypb.SecurityProfile{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateSecurityProfile(ctx, req)
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
								Usage: "delete security-profiles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "security_profile", Usage: "The security_profile.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s/locations/%s/securityProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile"))
									client, err := networksecurity.NewOrganizationSecurityProfileGroupClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteSecurityProfileRequest{Name: name}
									op, err := client.DeleteSecurityProfile(ctx, req)
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
						Name:  "server-tls-policies",
						Usage: "Manage server-tls-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list server-tls-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe server-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "server_tls_policy", Usage: "The server_tls_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("server_tls_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetServerTlsPolicyRequest{Name: name}
									resp, err := client.GetServerTlsPolicy(ctx, req)
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
								Usage: "create server-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "server-tls-policy-id", Usage: "The server tls policy id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "allow-open", Usage: "The allow open.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateServerTlsPolicyRequest{Parent: parent}
									req.ServerTlsPolicyId = cmd.String("server-tls-policy-id")
									req.ServerTlsPolicy = &networksecuritypb.ServerTlsPolicy{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										AllowOpen:   cmd.Bool("allow-open"),
									}
									op, err := client.CreateServerTlsPolicy(ctx, req)
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
								Usage: "update server-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "server_tls_policy", Usage: "The server_tls_policy.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.BoolFlag{Name: "allow-open", Usage: "The allow open.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("server_tls_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateServerTlsPolicyRequest{}
									req.ServerTlsPolicy = &networksecuritypb.ServerTlsPolicy{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
										AllowOpen:   cmd.Bool("allow-open"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("allow-open") {
										paths = append(paths, "allow_open")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateServerTlsPolicy(ctx, req)
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
								Usage: "delete server-tls-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "server_tls_policy", Usage: "The server_tls_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("server_tls_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteServerTlsPolicyRequest{Name: name}
									op, err := client.DeleteServerTlsPolicy(ctx, req)
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
						Name:  "tls-inspection-policies",
						Usage: "Manage tls-inspection-policies resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list tls-inspection-policies",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe tls-inspection-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls_inspection_policy", Usage: "The tls_inspection_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls_inspection_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetTlsInspectionPolicyRequest{Name: name}
									resp, err := client.GetTlsInspectionPolicy(ctx, req)
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
								Usage: "create tls-inspection-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls-inspection-policy-id", Usage: "The tls inspection policy id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "ca-pool", Usage: "The ca pool.", Required: true},
									&cli.StringFlag{Name: "trust-config", Usage: "The trust config.", Required: false},
									&cli.BoolFlag{Name: "exclude-public-ca-set", Usage: "The exclude public ca set.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateTlsInspectionPolicyRequest{Parent: parent}
									req.TlsInspectionPolicyId = cmd.String("tls-inspection-policy-id")
									req.TlsInspectionPolicy = &networksecuritypb.TlsInspectionPolicy{
										Name:               cmd.String("name"),
										Description:        cmd.String("description"),
										CaPool:             cmd.String("ca-pool"),
										TrustConfig:        cmd.String("trust-config"),
										ExcludePublicCaSet: cmd.Bool("exclude-public-ca-set"),
									}
									op, err := client.CreateTlsInspectionPolicy(ctx, req)
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
								Usage: "update tls-inspection-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls_inspection_policy", Usage: "The tls_inspection_policy.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "ca-pool", Usage: "The ca pool.", Required: false},
									&cli.StringFlag{Name: "trust-config", Usage: "The trust config.", Required: false},
									&cli.BoolFlag{Name: "exclude-public-ca-set", Usage: "The exclude public ca set.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls_inspection_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateTlsInspectionPolicyRequest{}
									req.TlsInspectionPolicy = &networksecuritypb.TlsInspectionPolicy{
										Name:               name,
										Name:               cmd.String("name"),
										Description:        cmd.String("description"),
										CaPool:             cmd.String("ca-pool"),
										TrustConfig:        cmd.String("trust-config"),
										ExcludePublicCaSet: cmd.Bool("exclude-public-ca-set"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("ca-pool") {
										paths = append(paths, "ca_pool")
									}
									if cmd.IsSet("trust-config") {
										paths = append(paths, "trust_config")
									}
									if cmd.IsSet("exclude-public-ca-set") {
										paths = append(paths, "exclude_public_ca_set")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateTlsInspectionPolicy(ctx, req)
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
								Usage: "delete tls-inspection-policies",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "tls_inspection_policy", Usage: "The tls_inspection_policy.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls_inspection_policy"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteTlsInspectionPolicyRequest{Name: name}
									op, err := client.DeleteTlsInspectionPolicy(ctx, req)
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
						Name:  "url-lists",
						Usage: "Manage url-lists resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list url-lists",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe url-lists",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "url_list", Usage: "The url_list.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("url_list"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.GetUrlListRequest{Name: name}
									resp, err := client.GetUrlList(ctx, req)
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
								Usage: "create url-lists",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "url-list-id", Usage: "The url list id.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.CreateUrlListRequest{Parent: parent}
									req.UrlListId = cmd.String("url-list-id")
									req.UrlList = &networksecuritypb.UrlList{
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									op, err := client.CreateUrlList(ctx, req)
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
								Usage: "update url-lists",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "url_list", Usage: "The url_list.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("url_list"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.UpdateUrlListRequest{}
									req.UrlList = &networksecuritypb.UrlList{
										Name:        name,
										Name:        cmd.String("name"),
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateUrlList(ctx, req)
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
								Usage: "delete url-lists",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "url_list", Usage: "The url_list.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("url_list"))
									client, err := networksecurity.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &networksecuritypb.DeleteUrlListRequest{Name: name}
									op, err := client.DeleteUrlList(ctx, req)
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
