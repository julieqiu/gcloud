// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package issuerswitch

import (
	issuerswitch "cloud.google.com/go/issuerswitch/apiv1"
	"cloud.google.com/go/issuerswitch/apiv1/issuerswitchpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the issuerswitch command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "issuerswitch",
		Usage: "manage Issuer switch API resources",
		Commands: []*cli.Command{
			{
				Name:  "complaint-transactions",
				Usage: "Manage complaint-transactions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list complaint-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of complaint transactions.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of transactions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListComplaintTransactions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ListComplaintTransactionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListComplaintTransactions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "export",
						Usage: "export complaint-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transaction-type", Usage: "Transaction type for the complaint transaction API.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ExportComplaintTransactionsRequest{
								Parent:          parent,
								TransactionType: issuerswitchpb.TransactionType(issuerswitchpb.TransactionType_value[cmd.String("transaction-type")]),
							}

							op, err := client.ExportComplaintTransactions(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "complaints",
				Usage: "Manage complaints resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create complaints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.CreateComplaintRequest{
								Parent: parent,
							}

							op, err := client.CreateComplaint(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "resolve",
						Usage: "resolve complaints",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "complaint", Usage: "The ID of the complaint.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "complaint.name" not yet supported.
							complaint_name := fmt.Sprintf("projects/%s/complaints/%s", cmd.String("project"), cmd.String("complaint"))
							fmt.Printf("Executing resolve on %s\n", complaint_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "disputes",
				Usage: "Manage disputes resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create disputes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.CreateDisputeRequest{
								Parent: parent,
							}

							op, err := client.CreateDispute(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "resolve",
						Usage: "resolve disputes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dispute", Usage: "The ID of the dispute.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "dispute.name" not yet supported.
							dispute_name := fmt.Sprintf("projects/%s/disputes/%s", cmd.String("project"), cmd.String("dispute"))
							fmt.Printf("Executing resolve on %s\n", dispute_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "financial-transactions",
				Usage: "Manage financial-transactions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list financial-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of financial transactions.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of transactions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListFinancialTransactions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ListFinancialTransactionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListFinancialTransactions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "export",
						Usage: "export financial-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transaction-type", Usage: "Transaction type for the financial transaction API.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ExportFinancialTransactionsRequest{
								Parent:          parent,
								TransactionType: issuerswitchpb.TransactionType(issuerswitchpb.TransactionType_value[cmd.String("transaction-type")]),
							}

							op, err := client.ExportFinancialTransactions(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "issuer-participants",
				Usage: "Manage issuer-participants resources",
				Commands: []*cli.Command{

					{
						Name:  "fetch",
						Usage: "fetch issuer-participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.FetchParticipantRequest{
								Parent: parent,
							}

							resp, err := client.FetchParticipant(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update issuer-participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.UpdateIssuerParticipantRequest{
								Parent: parent,
							}

							resp, err := client.UpdateIssuerParticipant(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "activate",
						Usage: "activate issuer-participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ParticipantStateChangeRequest{
								Parent: parent,
							}

							resp, err := client.ActivateParticipant(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "deactivate",
						Usage: "deactivate issuer-participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ParticipantStateChangeRequest{
								Parent: parent,
							}

							resp, err := client.DeactivateParticipant(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "mobile-number-changed",
						Usage: "mobile-number-changed issuer-participants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ParticipantStateChangeRequest{
								Parent: parent,
							}

							resp, err := client.MobileNumberChanged(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "mandate-transactions",
				Usage: "Manage mandate-transactions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list mandate-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of mandate transactions.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of transactions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListMandateTransactions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ListMandateTransactionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListMandateTransactions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "export",
						Usage: "export mandate-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "transaction-type", Usage: "Transaction type for the mandate transaction API.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ExportMandateTransactionsRequest{
								Parent:          parent,
								TransactionType: issuerswitchpb.TransactionType(issuerswitchpb.TransactionType_value[cmd.String("transaction-type")]),
							}

							op, err := client.ExportMandateTransactions(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},
				},
			},
			{
				Name:  "metadata",
				Usage: "Manage metadata resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list metadata",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of rule metadata to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListRuleMetadataRequest` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s", cmd.String("project"), cmd.String("rule"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ListRuleMetadataRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRuleMetadata(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "metadata-transactions",
				Usage: "Manage metadata-transactions resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list metadata-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "An expression that filters the list of metadata transactions.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of transactions to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListMetadataTransactions` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ListMetadataTransactionsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
								Filter:    cmd.String("filter"),
							}

							limit := cmd.Int("limit")
							it := client.ListMetadataTransactions(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "export",
						Usage: "export metadata-transactions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "api-type", Usage: "API type of the metadata transaction API.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ExportMetadataTransactionsRequest{
								Parent:  parent,
								ApiType: issuerswitchpb.ApiType(issuerswitchpb.ApiType_value[cmd.String("api-type")]),
							}

							op, err := client.ExportMetadataTransactions(ctx, req)
							if err != nil {
								return err
							}
							if cmd.Bool("async") {
								fmt.Println("Operation:", op.Name())
								return nil
							}
							resp, err := op.Wait(ctx)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
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
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							fmt.Printf("Executing describe on %s\n", name)
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
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of rules to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous `ListRulesRequest` call.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ListRulesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRules(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "values",
				Usage: "Manage values resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metadata", Usage: "The ID of the metadata.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of metadata values to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token received from a previous `ListRuleMetadataValuesRequest`.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.ListRuleMetadataValuesRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListRuleMetadataValues(ctx, req)
							count := 0
							for {
								resp, err := it.Next()
								if errors.Is(err, iterator.Done) {
									break
								}
								if err != nil {
									return err
								}
								out, err := runtime.FormatResponse(cmd.String("format"), resp)
								if err != nil {
									return err
								}
								fmt.Println(out)
								count++
								if limit > 0 && count >= limit {
									break
								}
							}
							return nil
						},
					},

					{
						Name:  "batch-create",
						Usage: "batch-create values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metadata", Usage: "The ID of the metadata.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.BatchCreateRuleMetadataValuesRequest{
								Parent: parent,
							}

							resp, err := client.BatchCreateRuleMetadataValues(ctx, req)
							if err != nil {
								return err
							}
							out, err := runtime.FormatResponse(cmd.String("format"), resp)
							if err != nil {
								return err
							}
							fmt.Println(out)
							return nil
						},
					},

					{
						Name:  "batch-delete",
						Usage: "batch-delete values",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "metadata", Usage: "The ID of the metadata.", Required: true},
							&cli.StringSliceFlag{Name: "names", Usage: "The names of the rule metadata values to delete.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "rule", Usage: "The ID of the rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute BatchDeleteRuleMetadataValues on %s?", parent)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := issuerswitch.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &issuerswitchpb.BatchDeleteRuleMetadataValuesRequest{
								Parent: parent,
								Names:  cmd.StringSlice("names"),
							}

							if err := client.BatchDeleteRuleMetadataValues(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},
				},
			},
		},
	}
}
