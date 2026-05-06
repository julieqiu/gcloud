package gkehub

import (
	"context"
	"fmt"
	"strings"

	gkehub "cloud.google.com/go/gkehub/apiv1"
	"cloud.google.com/go/gkehub/apiv1/gkehubpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud gkehub command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "gkehub",
		Usage: "manage GKE Hub resources",
		Commands: []*cli.Command{
			{
				Name:  "bindings",
				Usage: "Manage bindings resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
							&cli.StringFlag{Name: "membershipbinding", Usage: "The membershipbinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("membershipbinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetMembershipBindingRequest{Name: name}
							resp, err := client.GetMembershipBinding(ctx, req)
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
						Usage: "create bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
							&cli.StringFlag{Name: "membership-binding-id", Usage: "The membership binding id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateMembershipBindingRequest{Parent: parent}
							req.MembershipBindingId = cmd.String("membership-binding-id")
							req.MembershipBinding = &gkehubpb.MembershipBinding{
								Name: cmd.String("name"),
							}
							op, err := client.CreateMembershipBinding(ctx, req)
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
						Usage: "update bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
							&cli.StringFlag{Name: "membershipbinding", Usage: "The membershipbinding.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("membershipbinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateMembershipBindingRequest{}
							req.MembershipBinding = &gkehubpb.MembershipBinding{
								Name: name,
								Name: cmd.String("name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateMembershipBinding(ctx, req)
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
						Usage: "delete bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
							&cli.StringFlag{Name: "membershipbinding", Usage: "The membershipbinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("membershipbinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteMembershipBindingRequest{Name: name}
							op, err := client.DeleteMembershipBinding(ctx, req)
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
						Usage: "list bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListMembershipBindingsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListMembershipBindings(ctx, req)
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
				Name:  "features",
				Usage: "Manage features resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListFeaturesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListFeatures(ctx, req)
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
						Usage: "describe features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "feature", Usage: "The feature.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetFeatureRequest{Name: name}
							resp, err := client.GetFeature(ctx, req)
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
						Usage: "create features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "feature-id", Usage: "The feature id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateFeatureRequest{Parent: parent}
							req.FeatureId = cmd.String("feature-id")
							op, err := client.CreateFeature(ctx, req)
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
						Usage: "delete features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "feature", Usage: "The feature.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteFeatureRequest{Name: name}
							op, err := client.DeleteFeature(ctx, req)
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
						Usage: "update features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "feature", Usage: "The feature.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateFeatureRequest{}
							req.Resource = &gkehubpb.Feature{
								Name: name,
							}
							op, err := client.UpdateFeature(ctx, req)
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
				Name:  "fleets",
				Usage: "Manage fleets resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateFleetRequest{Parent: parent}
							req.Fleet = &gkehubpb.Fleet{
								DisplayName: cmd.String("display-name"),
							}
							op, err := client.CreateFleet(ctx, req)
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
						Usage: "describe fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "fleet", Usage: "The fleet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleets/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetFleetRequest{Name: name}
							resp, err := client.GetFleet(ctx, req)
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
						Usage: "update fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "fleet", Usage: "The fleet.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleets/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateFleetRequest{}
							req.Fleet = &gkehubpb.Fleet{
								Name:        name,
								DisplayName: cmd.String("display-name"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateFleet(ctx, req)
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
						Usage: "delete fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "fleet", Usage: "The fleet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleets/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteFleetRequest{Name: name}
							op, err := client.DeleteFleet(ctx, req)
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
						Usage: "list fleets",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
			{
				Name:  "memberships",
				Usage: "Manage memberships resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListMembershipsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListMemberships(ctx, req)
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
						Usage: "describe memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetMembershipRequest{Name: name}
							resp, err := client.GetMembership(ctx, req)
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
						Usage: "create memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership-id", Usage: "The membership id.", Required: true},
							&cli.StringFlag{Name: "external-id", Usage: "The external id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateMembershipRequest{Parent: parent}
							req.MembershipId = cmd.String("membership-id")
							req.Resource = &gkehubpb.Membership{
								ExternalId: cmd.String("external-id"),
							}
							op, err := client.CreateMembership(ctx, req)
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
						Usage: "delete memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteMembershipRequest{Name: name}
							op, err := client.DeleteMembership(ctx, req)
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
						Usage: "update memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
							&cli.StringFlag{Name: "external-id", Usage: "The external id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateMembershipRequest{}
							req.Resource = &gkehubpb.Membership{
								Name:       name,
								ExternalId: cmd.String("external-id"),
							}
							var paths []string
							if cmd.IsSet("external-id") {
								paths = append(paths, "external_id")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateMembership(ctx, req)
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
						Name:  "generate-connect-manifest",
						Usage: "generate-connect-manifest memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							fmt.Printf("Executing generate-connect-manifest on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "namespaces",
				Usage: "Manage namespaces resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("namespace"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetScopeNamespaceRequest{Name: name}
							resp, err := client.GetScopeNamespace(ctx, req)
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
						Usage: "create namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "scope-namespace-id", Usage: "The scope namespace id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateScopeNamespaceRequest{Parent: parent}
							req.ScopeNamespaceId = cmd.String("scope-namespace-id")
							req.ScopeNamespace = &gkehubpb.Namespace{
								Name:  cmd.String("name"),
								Scope: cmd.String("scope"),
							}
							op, err := client.CreateScopeNamespace(ctx, req)
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
						Usage: "update namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("namespace"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateScopeNamespaceRequest{}
							req.ScopeNamespace = &gkehubpb.Namespace{
								Name:  name,
								Name:  cmd.String("name"),
								Scope: cmd.String("scope"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("scope") {
								paths = append(paths, "scope")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateScopeNamespace(ctx, req)
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
						Usage: "delete namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "namespace", Usage: "The namespace.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("namespace"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteScopeNamespaceRequest{Name: name}
							op, err := client.DeleteScopeNamespace(ctx, req)
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
						Usage: "list namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListScopeNamespacesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListScopeNamespaces(ctx, req)
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
							client, err := gkehub.NewClient(ctx)
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
							client, err := gkehub.NewClient(ctx)
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
							client, err := gkehub.NewClient(ctx)
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
							client, err := gkehub.NewClient(ctx)
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
				Name:  "rbacrolebindings",
				Usage: "Manage rbacrolebindings resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The rbacrolebinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetScopeRBACRoleBindingRequest{Name: name}
							resp, err := client.GetScopeRBACRoleBinding(ctx, req)
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
						Usage: "create rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding-id", Usage: "The rbacrolebinding id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateScopeRBACRoleBindingRequest{Parent: parent}
							req.RbacrolebindingId = cmd.String("rbacrolebinding-id")
							req.Rbacrolebinding = &gkehubpb.RBACRoleBinding{
								Name: cmd.String("name"),
							}
							op, err := client.CreateScopeRBACRoleBinding(ctx, req)
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
						Usage: "update rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The rbacrolebinding.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateScopeRBACRoleBindingRequest{}
							req.Rbacrolebinding = &gkehubpb.RBACRoleBinding{
								Name: name,
								Name: cmd.String("name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateScopeRBACRoleBinding(ctx, req)
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
						Usage: "delete rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The rbacrolebinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteScopeRBACRoleBindingRequest{Name: name}
							op, err := client.DeleteScopeRBACRoleBinding(ctx, req)
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
						Usage: "list rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListScopeRBACRoleBindingsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListScopeRBACRoleBindings(ctx, req)
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
						Usage: "describe rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The rbacrolebinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetMembershipRBACRoleBindingRequest{Name: name}
							resp, err := client.GetMembershipRBACRoleBinding(ctx, req)
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
						Usage: "create rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding-id", Usage: "The rbacrolebinding id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateMembershipRBACRoleBindingRequest{Parent: parent}
							req.RbacrolebindingId = cmd.String("rbacrolebinding-id")
							req.Rbacrolebinding = &gkehubpb.RBACRoleBinding{
								Name: cmd.String("name"),
							}
							op, err := client.CreateMembershipRBACRoleBinding(ctx, req)
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
						Usage: "update rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The rbacrolebinding.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateMembershipRBACRoleBindingRequest{}
							req.Rbacrolebinding = &gkehubpb.RBACRoleBinding{
								Name: name,
								Name: cmd.String("name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateMembershipRBACRoleBinding(ctx, req)
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
						Usage: "delete rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "rbacrolebinding", Usage: "The rbacrolebinding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteMembershipRBACRoleBindingRequest{Name: name}
							op, err := client.DeleteMembershipRBACRoleBinding(ctx, req)
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
						Usage: "list rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListMembershipRBACRoleBindingsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListMembershipRBACRoleBindings(ctx, req)
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
						Name:  "generate-membership-rbac-role-binding-yaml",
						Usage: "generate-membership-rbac-role-binding-yaml rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing generate-membership-rbac-role-binding-yaml on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "scopes",
				Usage: "Manage scopes resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list scopes",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.GetScopeRequest{Name: name}
							resp, err := client.GetScope(ctx, req)
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
						Usage: "create scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope-id", Usage: "The scope id.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.CreateScopeRequest{Parent: parent}
							req.ScopeId = cmd.String("scope-id")
							req.Scope = &gkehubpb.Scope{
								Name: cmd.String("name"),
							}
							op, err := client.CreateScope(ctx, req)
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
						Usage: "update scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.UpdateScopeRequest{}
							req.Scope = &gkehubpb.Scope{
								Name: name,
								Name: cmd.String("name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							op, err := client.UpdateScope(ctx, req)
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
						Usage: "delete scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &gkehubpb.DeleteScopeRequest{Name: name}
							op, err := client.DeleteScope(ctx, req)
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
						Usage: "list scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListScopesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListScopes(ctx, req)
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
						Usage: "list scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := gkehub.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &gkehubpb.ListPermittedScopesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListPermittedScopes(ctx, req)
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
		},
	}
}
