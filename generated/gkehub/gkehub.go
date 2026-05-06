package gkehub

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							fmt.Printf("Executing create on %s\n", parent)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s/bindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"), cmd.String("membershipbinding"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							fmt.Printf("Executing list on %s\n", parent)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create features",
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
						Name:  "delete",
						Usage: "delete features",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "feature", Usage: "The feature.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/features/%s", cmd.String("project"), cmd.String("location"), cmd.String("feature"))
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing update on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update fleets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "fleet", Usage: "The fleet.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/fleets/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create memberships",
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
						Name:  "delete",
						Usage: "delete memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update memberships",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "membership", Usage: "The membership.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/memberships/%s", cmd.String("project"), cmd.String("location"), cmd.String("membership"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing create on %s\n", parent)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/namespaces/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("namespace"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list namespaces",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing list on %s\n", parent)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing cancel on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing create on %s\n", parent)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing create on %s\n", parent)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s/rbacrolebindings/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"), cmd.String("rbacrolebinding"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list rbacrolebindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create scopes",
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
						Usage: "update scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "scope", Usage: "The scope.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/scopes/%s", cmd.String("project"), cmd.String("location"), cmd.String("scope"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list scopes",
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
						Name:  "list",
						Usage: "list scopes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
				},
			},
		},
	}
}
