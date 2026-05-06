package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
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
				Name:  "cloudresourcemanager",
				Usage: "manage Cloud Resource Manager API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "effective-tags",
						Usage: "Manage effective-tags resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list effective-tags",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
						},
					},
					{
						Name:  "folders",
						Usage: "Manage folders resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("folders/%s", cmd.String("folder"))
									client, err := resourcemanager.NewFoldersClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.GetFolderRequest{Name: name}
									resp, err := client.GetFolder(ctx, req)
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
								Usage: "list folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "search",
								Usage: "search folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
									&cli.StringFlag{Name: "parent", Usage: "The parent.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("folders/%s", cmd.String("folder"))
									client, err := resourcemanager.NewFoldersClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.UpdateFolderRequest{}
									req.Folder = &resourcemanagerpb.Folder{
										Name:        name,
										Parent:      cmd.String("parent"),
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("parent") {
										paths = append(paths, "parent")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateFolder(ctx, req)
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
								Name:  "move",
								Usage: "move folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
									&cli.StringFlag{Name: "destination-parent", Usage: "The destination parent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("folders/%s", cmd.String("folder"))
									client, err := resourcemanager.NewFoldersClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.MoveFolderRequest{Name: name}
									req.DestinationParent = cmd.String("destination-parent")
									op, err := client.MoveFolder(ctx, req)
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
								Usage: "delete folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("folders/%s", cmd.String("folder"))
									client, err := resourcemanager.NewFoldersClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.DeleteFolderRequest{Name: name}
									op, err := client.DeleteFolder(ctx, req)
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
								Name:  "undelete",
								Usage: "undelete folders",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "folder", Usage: "The folder.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("folders/%s", cmd.String("folder"))
									client, err := resourcemanager.NewFoldersClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.UndeleteFolderRequest{Name: name}
									op, err := client.UndeleteFolder(ctx, req)
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
								Name:  "get-iam-policy",
								Usage: "get-iam-policy folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions folders",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "namespaced",
						Usage: "Manage namespaced resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe namespaced",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_key", Usage: "The tag_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagKeys/%s", cmd.String("tag_key"))
									client, err := resourcemanager.NewTagKeysClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.GetNamespacedTagKeyRequest{Name: name}
									resp, err := client.GetNamespacedTagKey(ctx, req)
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
								Usage: "describe namespaced",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_value", Usage: "The tag_value.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagValues/%s", cmd.String("tag_value"))
									client, err := resourcemanager.NewTagValuesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.GetNamespacedTagValueRequest{Name: name}
									resp, err := client.GetNamespacedTagValue(ctx, req)
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
								Name:  "describe",
								Usage: "describe operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
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
								Name:  "describe",
								Usage: "describe operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
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
								Name:  "describe",
								Usage: "describe operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
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
								Name:  "describe",
								Usage: "describe operations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing describe...")
									return nil
								},
							},
						},
					},
					{
						Name:  "organizations",
						Usage: "Manage organizations resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe organizations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("organizations/%s", cmd.String("organization"))
									client, err := resourcemanager.NewOrganizationsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.GetOrganizationRequest{Name: name}
									resp, err := client.GetOrganization(ctx, req)
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
								Name:  "search",
								Usage: "search organizations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search...")
									return nil
								},
							},
							{
								Name:  "get-iam-policy",
								Usage: "get-iam-policy organizations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy organizations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions organizations",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
									name := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := resourcemanager.NewProjectsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.GetProjectRequest{Name: name}
									resp, err := client.GetProject(ctx, req)
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
								Usage: "list projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "search",
								Usage: "search projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing search...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "parent", Usage: "The parent.", Required: false},
									&cli.StringFlag{Name: "project-id", Usage: "The project id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := resourcemanager.NewProjectsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.UpdateProjectRequest{}
									req.Project = &resourcemanagerpb.Project{
										Name:        name,
										Parent:      cmd.String("parent"),
										ProjectId:   cmd.String("project-id"),
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("parent") {
										paths = append(paths, "parent")
									}
									if cmd.IsSet("project-id") {
										paths = append(paths, "project_id")
									}
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateProject(ctx, req)
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
								Name:  "move",
								Usage: "move projects",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "destination-parent", Usage: "The destination parent.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := resourcemanager.NewProjectsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.MoveProjectRequest{Name: name}
									req.DestinationParent = cmd.String("destination-parent")
									op, err := client.MoveProject(ctx, req)
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
								Usage: "delete projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := resourcemanager.NewProjectsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.DeleteProjectRequest{Name: name}
									op, err := client.DeleteProject(ctx, req)
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
								Name:  "undelete",
								Usage: "undelete projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := resourcemanager.NewProjectsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.UndeleteProjectRequest{Name: name}
									op, err := client.UndeleteProject(ctx, req)
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
								Name:  "get-iam-policy",
								Usage: "get-iam-policy projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions projects",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "tag-bindings",
						Usage: "Manage tag-bindings resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list tag-bindings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "create",
								Usage: "create tag-bindings",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "delete",
								Usage: "delete tag-bindings",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_binding", Usage: "The tag_binding.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagBindings/%s", cmd.String("tag_binding"))
									client, err := resourcemanager.NewTagBindingsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.DeleteTagBindingRequest{Name: name}
									op, err := client.DeleteTagBinding(ctx, req)
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
						Name:  "tag-holds",
						Usage: "Manage tag-holds resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create tag-holds",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_value", Usage: "The tag_value.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "holder", Usage: "The holder.", Required: true},
									&cli.StringFlag{Name: "origin", Usage: "The origin.", Required: false},
									&cli.StringFlag{Name: "help-link", Usage: "The help link.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("tagValues/%s", cmd.String("tag_value"))
									client, err := resourcemanager.NewTagHoldsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.CreateTagHoldRequest{Parent: parent}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.TagHold = &resourcemanagerpb.TagHold{
										Holder:   cmd.String("holder"),
										Origin:   cmd.String("origin"),
										HelpLink: cmd.String("help-link"),
									}
									op, err := client.CreateTagHold(ctx, req)
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
								Usage: "delete tag-holds",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_value", Usage: "The tag_value.", Required: true},
									&cli.StringFlag{Name: "tag_hold", Usage: "The tag_hold.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagValues/%s/tagHolds/%s", cmd.String("tag_value"), cmd.String("tag_hold"))
									client, err := resourcemanager.NewTagHoldsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.DeleteTagHoldRequest{Name: name}
									op, err := client.DeleteTagHold(ctx, req)
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
								Usage: "list tag-holds",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_value", Usage: "The tag_value.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("tagValues/%s", cmd.String("tag_value"))
									client, err := resourcemanager.NewTagHoldsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &resourcemanagerpb.ListTagHoldsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListTagHolds(ctx, req)
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
						Name:  "tag-keys",
						Usage: "Manage tag-keys resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list tag-keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe tag-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_key", Usage: "The tag_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagKeys/%s", cmd.String("tag_key"))
									client, err := resourcemanager.NewTagKeysClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.GetTagKeyRequest{Name: name}
									resp, err := client.GetTagKey(ctx, req)
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
								Usage: "create tag-keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update tag-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_key", Usage: "The tag_key.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "parent", Usage: "The parent.", Required: false},
									&cli.StringFlag{Name: "short-name", Usage: "The short name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagKeys/%s", cmd.String("tag_key"))
									client, err := resourcemanager.NewTagKeysClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.UpdateTagKeyRequest{}
									req.TagKey = &resourcemanagerpb.TagKey{
										Name:        name,
										Name:        cmd.String("name"),
										Parent:      cmd.String("parent"),
										ShortName:   cmd.String("short-name"),
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("parent") {
										paths = append(paths, "parent")
									}
									if cmd.IsSet("short-name") {
										paths = append(paths, "short_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateTagKey(ctx, req)
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
								Usage: "delete tag-keys",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_key", Usage: "The tag_key.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagKeys/%s", cmd.String("tag_key"))
									client, err := resourcemanager.NewTagKeysClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.DeleteTagKeyRequest{Name: name}
									op, err := client.DeleteTagKey(ctx, req)
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
								Name:  "get-iam-policy",
								Usage: "get-iam-policy tag-keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy tag-keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions tag-keys",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
									return nil
								},
							},
						},
					},
					{
						Name:  "tag-values",
						Usage: "Manage tag-values resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list tag-values",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe tag-values",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_value", Usage: "The tag_value.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagValues/%s", cmd.String("tag_value"))
									client, err := resourcemanager.NewTagValuesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.GetTagValueRequest{Name: name}
									resp, err := client.GetTagValue(ctx, req)
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
								Usage: "create tag-values",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing create...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update tag-values",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_value", Usage: "The tag_value.", Required: true},
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
									&cli.StringFlag{Name: "parent", Usage: "The parent.", Required: false},
									&cli.StringFlag{Name: "short-name", Usage: "The short name.", Required: false},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagValues/%s", cmd.String("tag_value"))
									client, err := resourcemanager.NewTagValuesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.UpdateTagValueRequest{}
									req.TagValue = &resourcemanagerpb.TagValue{
										Name:        name,
										Name:        cmd.String("name"),
										Parent:      cmd.String("parent"),
										ShortName:   cmd.String("short-name"),
										Description: cmd.String("description"),
										Etag:        cmd.String("etag"),
									}
									var paths []string
									if cmd.IsSet("name") {
										paths = append(paths, "name")
									}
									if cmd.IsSet("parent") {
										paths = append(paths, "parent")
									}
									if cmd.IsSet("short-name") {
										paths = append(paths, "short_name")
									}
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									if cmd.IsSet("etag") {
										paths = append(paths, "etag")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateTagValue(ctx, req)
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
								Usage: "delete tag-values",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "tag_value", Usage: "The tag_value.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("tagValues/%s", cmd.String("tag_value"))
									client, err := resourcemanager.NewTagValuesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &resourcemanagerpb.DeleteTagValueRequest{Name: name}
									op, err := client.DeleteTagValue(ctx, req)
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
								Name:  "get-iam-policy",
								Usage: "get-iam-policy tag-values",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing get-iam-policy...")
									return nil
								},
							},
							{
								Name:  "set-iam-policy",
								Usage: "set-iam-policy tag-values",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing set-iam-policy...")
									return nil
								},
							},
							{
								Name:  "test-iam-permissions",
								Usage: "test-iam-permissions tag-values",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing test-iam-permissions...")
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
