package iam

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud iam command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "iam",
		Usage: "manage Identity and Access Management (IAM) API resources",
		Commands: []*cli.Command{
			{
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{
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
				},
			},
			{
				Name:  "policy-bindings",
				Usage: "Manage policy-bindings resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "policy_binding", Usage: "The policy_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/policyBindings/%s", cmd.String("organization"), cmd.String("location"), cmd.String("policy_binding"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "policy_binding", Usage: "The policy_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/policyBindings/%s", cmd.String("organization"), cmd.String("location"), cmd.String("policy_binding"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "policy_binding", Usage: "The policy_binding.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/policyBindings/%s", cmd.String("organization"), cmd.String("location"), cmd.String("policy_binding"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list policy-bindings",
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
						Name:  "search-target-policy-bindings",
						Usage: "search-target-policy-bindings policy-bindings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing search-target-policy-bindings on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "principal-access-boundary-policies",
				Usage: "Manage principal-access-boundary-policies resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "principal_access_boundary_policy", Usage: "The principal_access_boundary_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal_access_boundary_policy"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "principal_access_boundary_policy", Usage: "The principal_access_boundary_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal_access_boundary_policy"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "principal_access_boundary_policy", Usage: "The principal_access_boundary_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal_access_boundary_policy"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list principal-access-boundary-policies",
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
						Name:  "search-policy-bindings",
						Usage: "search-policy-bindings principal-access-boundary-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "principal_access_boundary_policy", Usage: "The principal_access_boundary_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/principalAccessBoundaryPolicies/%s", cmd.String("organization"), cmd.String("location"), cmd.String("principal_access_boundary_policy"))
							fmt.Printf("Executing search-policy-bindings on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
