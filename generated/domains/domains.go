package domains

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud domains command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "domains",
		Usage: "manage Cloud Domains API resources",
		Commands: []*cli.Command{
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
				},
			},
			{
				Name:  "registrations",
				Usage: "Manage registrations resources",
				Commands: []*cli.Command{
					{
						Name:  "search-domains",
						Usage: "search-domains registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search-domains...")
							return nil
						},
					},
					{
						Name:  "retrieve-register-parameters",
						Usage: "retrieve-register-parameters registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing retrieve-register-parameters...")
							return nil
						},
					},
					{
						Name:  "register",
						Usage: "register registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing register on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "retrieve-transfer-parameters",
						Usage: "retrieve-transfer-parameters registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing retrieve-transfer-parameters...")
							return nil
						},
					},
					{
						Name:  "transfer",
						Usage: "transfer registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing transfer on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing update on %s\n", name)
							return nil
						},
					},
					{
						Name:  "configure-management-settings",
						Usage: "configure-management-settings registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing configure-management-settings...")
							return nil
						},
					},
					{
						Name:  "configure-dns-settings",
						Usage: "configure-dns-settings registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing configure-dns-settings...")
							return nil
						},
					},
					{
						Name:  "configure-contact-settings",
						Usage: "configure-contact-settings registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing configure-contact-settings...")
							return nil
						},
					},
					{
						Name:  "export",
						Usage: "export registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing export on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete registrations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "registration", Usage: "The registration.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/registrations/%s", cmd.String("project"), cmd.String("location"), cmd.String("registration"))
							fmt.Printf("Executing delete on %s\n", name)
							return nil
						},
					},
					{
						Name:  "retrieve-authorization-code",
						Usage: "retrieve-authorization-code registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing retrieve-authorization-code...")
							return nil
						},
					},
					{
						Name:  "reset-authorization-code",
						Usage: "reset-authorization-code registrations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing reset-authorization-code...")
							return nil
						},
					},
				},
			},
		},
	}
}
