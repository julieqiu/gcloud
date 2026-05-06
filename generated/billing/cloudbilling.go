package cloudbilling

import (
	"context"
	"fmt"

	billing "cloud.google.com/go/billing/apiv1"
	"cloud.google.com/go/billing/apiv1/billingpb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud cloudbilling command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudbilling",
		Usage: "manage Cloud Billing API resources",
		Commands: []*cli.Command{
			{
				Name:  "billing-accounts",
				Usage: "Manage billing-accounts resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing_account", Usage: "The billing_account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billing_account"))
							client, err := billing.NewCloudBillingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &billingpb.GetBillingAccountRequest{Name: name}
							resp, err := client.GetBillingAccount(ctx, req)
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
						Usage: "list billing-accounts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing_account", Usage: "The billing_account.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "master-billing-account", Usage: "The master billing account.", Required: false},
							&cli.StringFlag{Name: "currency-code", Usage: "The currency code.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billing_account"))
							client, err := billing.NewCloudBillingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &billingpb.UpdateBillingAccountRequest{}
							req.Account = &billingpb.BillingAccount{
								Name:                 name,
								DisplayName:          cmd.String("display-name"),
								MasterBillingAccount: cmd.String("master-billing-account"),
								CurrencyCode:         cmd.String("currency-code"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("master-billing-account") {
								paths = append(paths, "master_billing_account")
							}
							if cmd.IsSet("currency-code") {
								paths = append(paths, "currency_code")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateBillingAccount(ctx, req)
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
						Usage: "create billing-accounts",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing create...")
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billingAccount", Usage: "The billingAccount.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billingAccount"))
							client, err := billing.NewCloudBillingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.GetIamPolicyRequest{Resource: name}
							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billingAccount", Usage: "The billingAccount.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billingAccount"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billingAccount", Usage: "The billingAccount.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billingAccount"))
							client, err := billing.NewCloudBillingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.TestIamPermissionsRequest{Resource: name}
							req.Permissions = cmd.StringSlice("permissions")
							resp, err := client.TestIamPermissions(ctx, req)
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
						Usage: "move billing-accounts",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing_account", Usage: "The billing_account.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("billingAccounts/%s", cmd.String("billing_account"))
							fmt.Printf("Executing move on %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "billing-info",
				Usage: "Manage billing-info resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe billing-info",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := billing.NewCloudBillingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &billingpb.GetProjectBillingInfoRequest{Name: name}
							resp, err := client.GetProjectBillingInfo(ctx, req)
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
						Usage: "update billing-info",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "billing-account-name", Usage: "The billing account name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/billingInfo", cmd.String("project"))
							client, err := billing.NewCloudBillingClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &billingpb.UpdateProjectBillingInfoRequest{}
							req.ProjectBillingInfo = &billingpb.ProjectBillingInfo{
								Name:               name,
								BillingAccountName: cmd.String("billing-account-name"),
							}
							var paths []string
							if cmd.IsSet("billing-account-name") {
								paths = append(paths, "billing_account_name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateProjectBillingInfo(ctx, req)
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
				Name:  "projects",
				Usage: "Manage projects resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list projects",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
			{
				Name:  "services",
				Usage: "Manage services resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list services",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
			{
				Name:  "skus",
				Usage: "Manage skus resources",
				Commands: []*cli.Command{
					{
						Name:  "list",
						Usage: "list skus",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
		},
	}
}
