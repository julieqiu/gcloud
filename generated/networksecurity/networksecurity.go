package networksecurity

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud networksecurity command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "networksecurity",
		Usage: "manage Network Security API resources",
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create address-groups",
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
						Usage: "update address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create address-groups",
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
						Usage: "update address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "address_group", Usage: "The address_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("address_group"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
							return nil
						},
					},
					{
						Name:  "set-iam-policy",
						Usage: "set-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing get-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions address-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "addressGroup", Usage: "The addressGroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/addressGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("addressGroup"))
							fmt.Printf("Executing test-iam-permissions on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create authorization-policies",
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
						Usage: "update authorization-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "authorization_policy", Usage: "The authorization_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authorizationPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authorization_policy"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create authz-policies",
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
						Usage: "update authz-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "authz_policy", Usage: "The authz_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/authzPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("authz_policy"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create backend-authentication-configs",
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
						Usage: "update backend-authentication-configs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "backend_authentication_config", Usage: "The backend_authentication_config.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/backendAuthenticationConfigs/%s", cmd.String("project"), cmd.String("location"), cmd.String("backend_authentication_config"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create client-tls-policies",
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
						Usage: "update client-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "client_tls_policy", Usage: "The client_tls_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/clientTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("client_tls_policy"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create dns-threat-detectors",
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
						Usage: "update dns-threat-detectors",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "dns_threat_detector", Usage: "The dns_threat_detector.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/dnsThreatDetectors/%s", cmd.String("project"), cmd.String("location"), cmd.String("dns_threat_detector"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create firewall-endpoint-associations",
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
						Usage: "delete firewall-endpoint-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "firewall_endpoint_association", Usage: "The firewall_endpoint_association.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall_endpoint_association"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update firewall-endpoint-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "firewall_endpoint_association", Usage: "The firewall_endpoint_association.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/firewallEndpointAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("firewall_endpoint_association"))
							fmt.Printf("Executing update on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create firewall-endpoints",
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
						Name:  "delete",
						Usage: "delete firewall-endpoints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "firewall_endpoint", Usage: "The firewall_endpoint.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall_endpoint"))
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/firewallEndpoints/%s", cmd.String("organization"), cmd.String("location"), cmd.String("firewall_endpoint"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create gateway-security-policies",
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
						Usage: "update gateway-security-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create intercept-deployment-groups",
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
						Usage: "update intercept-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "intercept_deployment_group", Usage: "The intercept_deployment_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment_group"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create intercept-deployments",
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
						Usage: "update intercept-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "intercept_deployment", Usage: "The intercept_deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_deployment"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create intercept-endpoint-group-associations",
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
						Usage: "update intercept-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "intercept_endpoint_group_association", Usage: "The intercept_endpoint_group_association.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group_association"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create intercept-endpoint-groups",
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
						Usage: "update intercept-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "intercept_endpoint_group", Usage: "The intercept_endpoint_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/interceptEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("intercept_endpoint_group"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create mirroring-deployment-groups",
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
						Usage: "update mirroring-deployment-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "mirroring_deployment_group", Usage: "The mirroring_deployment_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeploymentGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment_group"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create mirroring-deployments",
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
						Usage: "update mirroring-deployments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "mirroring_deployment", Usage: "The mirroring_deployment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringDeployments/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_deployment"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create mirroring-endpoint-group-associations",
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
						Usage: "update mirroring-endpoint-group-associations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "mirroring_endpoint_group_association", Usage: "The mirroring_endpoint_group_association.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroupAssociations/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group_association"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create mirroring-endpoint-groups",
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
						Usage: "update mirroring-endpoint-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "mirroring_endpoint_group", Usage: "The mirroring_endpoint_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/mirroringEndpointGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("mirroring_endpoint_group"))
							fmt.Printf("Executing update on %s\n", name)
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
				Name:  "rules",
				Usage: "Manage rules resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list rules",
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
						Usage: "describe rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"), cmd.String("rule"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "gateway_security_policy", Usage: "The gateway_security_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"))
							fmt.Printf("Executing create on %s\n", parent)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/gatewaySecurityPolicies/%s/rules/%s", cmd.String("project"), cmd.String("location"), cmd.String("gateway_security_policy"), cmd.String("rule"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create security-profile-groups",
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
						Name:  "update",
						Usage: "update security-profile-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "security_profile_group", Usage: "The security_profile_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/securityProfileGroups/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile_group"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("organizations/%s/locations/%s", cmd.String("organization"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", parent)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create security-profiles",
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
						Name:  "update",
						Usage: "update security-profiles",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "organization", Usage: "The organization.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "security_profile", Usage: "The security_profile.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("organizations/%s/locations/%s/securityProfiles/%s", cmd.String("organization"), cmd.String("location"), cmd.String("security_profile"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create server-tls-policies",
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
						Usage: "update server-tls-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "server_tls_policy", Usage: "The server_tls_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("server_tls_policy"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create tls-inspection-policies",
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
						Usage: "update tls-inspection-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "tls_inspection_policy", Usage: "The tls_inspection_policy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/tlsInspectionPolicies/%s", cmd.String("project"), cmd.String("location"), cmd.String("tls_inspection_policy"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
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
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create url-lists",
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
						Usage: "update url-lists",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "url_list", Usage: "The url_list.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/urlLists/%s", cmd.String("project"), cmd.String("location"), cmd.String("url_list"))
							fmt.Printf("Executing update on %s\n", name)
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
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
