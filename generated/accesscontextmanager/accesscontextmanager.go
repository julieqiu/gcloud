package accesscontextmanager

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud accesscontextmanager command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "accesscontextmanager",
		Usage: "manage Access Context Manager API resources",
		Commands: []*cli.Command{
			{
				Name:  "access-levels",
				Usage: "Manage access-levels resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
							&cli.StringFlag{Name: "access_level", Usage: "The access_level.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/accessLevels/%s", cmd.String("access_policy"), cmd.String("access_level"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
							&cli.StringFlag{Name: "access_level", Usage: "The access_level.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/accessLevels/%s", cmd.String("access_policy"), cmd.String("access_level"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
							&cli.StringFlag{Name: "access_level", Usage: "The access_level.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/accessLevels/%s", cmd.String("access_policy"), cmd.String("access_level"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "replace-all",
						Usage: "replace-all access-levels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing replace-all on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "access-policies",
				Usage: "Manage access-policies resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list access-policies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create access-policies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "accessPolicie", Usage: "The accessPolicie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("accessPolicie"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "accessPolicie", Usage: "The accessPolicie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("accessPolicie"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "accessPolicie", Usage: "The accessPolicie.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s", cmd.String("accessPolicie"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "gcp-user-access-bindings",
				Usage: "Manage gcp-user-access-bindings resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list gcp-user-access-bindings",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "gcp_user_access_binding", Usage: "The gcp_user_access_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/%s", cmd.String("organization"), cmd.String("gcp_user_access_binding"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create gcp-user-access-bindings",
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
						Usage: "update gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "gcp_user_access_binding", Usage: "The gcp_user_access_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/%s", cmd.String("organization"), cmd.String("gcp_user_access_binding"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete gcp-user-access-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "gcp_user_access_binding", Usage: "The gcp_user_access_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/gcpUserAccessBindings/%s", cmd.String("organization"), cmd.String("gcp_user_access_binding"))
							fmt.Printf("Executing delete on %s\n", name)
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
				},
			},
			{
				Name:  "service-perimeters",
				Usage: "Manage service-perimeters resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
							&cli.StringFlag{Name: "service_perimeter", Usage: "The service_perimeter.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", cmd.String("access_policy"), cmd.String("service_perimeter"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
							&cli.StringFlag{Name: "service_perimeter", Usage: "The service_perimeter.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", cmd.String("access_policy"), cmd.String("service_perimeter"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
							&cli.StringFlag{Name: "service_perimeter", Usage: "The service_perimeter.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", cmd.String("access_policy"), cmd.String("service_perimeter"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "replace-all",
						Usage: "replace-all service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing replace-all on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "commit",
						Usage: "commit service-perimeters",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "access_policy", Usage: "The access_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("accessPolicies/%s", cmd.String("access_policy"))
							fmt.Printf("Executing commit on %s\n", parent)
							return nil
						},
					},
				},
			},
		},
	}
}
