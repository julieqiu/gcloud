package recaptchaenterprise

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud recaptchaenterprise command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "recaptchaenterprise",
		Usage: "manage reCAPTCHA Enterprise API resources",
		Commands: []*cli.Command{
			{
				Name:  "assessments",
				Usage: "Manage assessments resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create assessments",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "annotate",
						Usage: "annotate assessments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "assessment", Usage: "The assessment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/assessments/%s", cmd.String("project"), cmd.String("assessment"))
							fmt.Printf("Executing annotate on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "firewallpolicies",
				Usage: "Manage firewallpolicies resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create firewallpolicies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list firewallpolicies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewallpolicy", Usage: "The firewallpolicy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicy"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewallpolicy", Usage: "The firewallpolicy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicy"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete firewallpolicies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "firewallpolicy", Usage: "The firewallpolicy.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/firewallpolicies/%s", cmd.String("project"), cmd.String("firewallpolicy"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "reorder",
						Usage: "reorder firewallpolicies",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing reorder...")
							return nil
						},
					},
				},
			},
			{
				Name:  "keys",
				Usage: "Manage keys resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create keys",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list keys",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "retrieve-legacy-secret-key",
						Usage: "retrieve-legacy-secret-key keys",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing retrieve-legacy-secret-key...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "migrate",
						Usage: "migrate keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing migrate on %s\n", name)
							return nil
						},
					},
					{
						Name:  "add-ip-override",
						Usage: "add-ip-override keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing add-ip-override on %s\n", name)
							return nil
						},
					},
					{
						Name:  "remove-ip-override",
						Usage: "remove-ip-override keys",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing remove-ip-override on %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list keys",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
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
							&cli.StringFlag{Name: "relatedaccountgroup", Usage: "The relatedaccountgroup.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/relatedaccountgroups/%s", cmd.String("project"), cmd.String("relatedaccountgroup"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
				},
			},
			{
				Name:  "metrics",
				Usage: "Manage metrics resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe metrics",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "key", Usage: "The key.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/keys/%s/metrics", cmd.String("project"), cmd.String("key"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "relatedaccountgroupmemberships",
				Usage: "Manage relatedaccountgroupmemberships resources",
				Commands: []*cli.Command{
					{
						Name:  "search",
						Usage: "search relatedaccountgroupmemberships",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search...")
							return nil
						},
					},
				},
			},
			{
				Name:  "relatedaccountgroups",
				Usage: "Manage relatedaccountgroups resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list relatedaccountgroups",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", parent)
							return nil
						},
					},
				},
			},
		},
	}
}
