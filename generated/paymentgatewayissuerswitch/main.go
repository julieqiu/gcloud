package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	issuerswitch "cloud.google.com/go/paymentgateway/issuerswitch/apiv1"
	"cloud.google.com/go/paymentgateway/issuerswitch/apiv1/issuerswitchpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "issuerswitch",
				Usage: "manage Issuer switch API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "complaint-transactions",
						Usage: "Manage complaint-transactions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list complaint-transactions",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchTransactionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &issuerswitchpb.ListComplaintTransactionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListComplaintTransactions(ctx, req)
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
								Name:  "export",
								Usage: "export complaint-transactions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing export...")
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
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchResolutionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &issuerswitchpb.CreateComplaintRequest{Parent: parent}
									req.Complaint = &issuerswitchpb.Complaint{
										Name: cmd.String("name"),
									}
									op, err := client.CreateComplaint(ctx, req)
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
								Name:  "resolve",
								Usage: "resolve complaints",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "complaint", Usage: "The complaint.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/complaints/%s", cmd.String("project"), cmd.String("complaint"))
									client, err := issuerswitch.NewIssuerSwitchResolutionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &issuerswitchpb.ResolveComplaintRequest{Name: name}
									op, err := client.ResolveComplaint(ctx, req)
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
						Name:  "disputes",
						Usage: "Manage disputes resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create disputes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchResolutionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &issuerswitchpb.CreateDisputeRequest{Parent: parent}
									req.Dispute = &issuerswitchpb.Dispute{
										Name: cmd.String("name"),
									}
									op, err := client.CreateDispute(ctx, req)
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
								Name:  "resolve",
								Usage: "resolve disputes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "dispute", Usage: "The dispute.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/disputes/%s", cmd.String("project"), cmd.String("dispute"))
									client, err := issuerswitch.NewIssuerSwitchResolutionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &issuerswitchpb.ResolveDisputeRequest{Name: name}
									op, err := client.ResolveDispute(ctx, req)
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
						Name:  "financial-transactions",
						Usage: "Manage financial-transactions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list financial-transactions",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchTransactionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &issuerswitchpb.ListFinancialTransactionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListFinancialTransactions(ctx, req)
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
								Name:  "export",
								Usage: "export financial-transactions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing export...")
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
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing fetch...")
									return nil
								},
							},
							{
								Name:  "update",
								Usage: "update issuer-participants",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing update...")
									return nil
								},
							},
							{
								Name:  "activate",
								Usage: "activate issuer-participants",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing activate...")
									return nil
								},
							},
							{
								Name:  "deactivate",
								Usage: "deactivate issuer-participants",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing deactivate...")
									return nil
								},
							},
							{
								Name:  "mobile-number-changed",
								Usage: "mobile-number-changed issuer-participants",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing mobile-number-changed...")
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchTransactionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &issuerswitchpb.ListMandateTransactionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMandateTransactions(ctx, req)
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
								Name:  "export",
								Usage: "export mandate-transactions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing export...")
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
									&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/rules/%s", cmd.String("project"), cmd.String("rule"))
									client, err := issuerswitch.NewIssuerSwitchRulesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &issuerswitchpb.ListRuleMetadataRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRuleMetadata(ctx, req)
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
						Name:  "metadata-transactions",
						Usage: "Manage metadata-transactions resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list metadata-transactions",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchTransactionsClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &issuerswitchpb.ListMetadataTransactionsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListMetadataTransactions(ctx, req)
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
								Name:  "export",
								Usage: "export metadata-transactions",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing export...")
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchParticipantsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := issuerswitch.NewIssuerSwitchParticipantsClient(ctx)
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
								Name:  "list",
								Usage: "list operations",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchResolutionsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := issuerswitch.NewIssuerSwitchResolutionsClient(ctx)
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
								Name:  "list",
								Usage: "list operations",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchRulesClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := issuerswitch.NewIssuerSwitchRulesClient(ctx)
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
								Name:  "list",
								Usage: "list operations",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchTransactionsClient(ctx)
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
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
									client, err := issuerswitch.NewIssuerSwitchTransactionsClient(ctx)
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
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := issuerswitch.NewIssuerSwitchRulesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &issuerswitchpb.ListRulesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRules(ctx, req)
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
						Name:  "values",
						Usage: "Manage values resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list values",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
									&cli.StringFlag{Name: "metadata", Usage: "The metadata.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
									client, err := issuerswitch.NewIssuerSwitchRulesClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &issuerswitchpb.ListRuleMetadataValuesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRuleMetadataValues(ctx, req)
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
								Name:  "batch-create",
								Usage: "batch-create values",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
									&cli.StringFlag{Name: "metadata", Usage: "The metadata.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
									fmt.Printf("Executing batch-create on %s\n", parent)
									return nil
								},
							},
							{
								Name:  "batch-delete",
								Usage: "batch-delete values",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "rule", Usage: "The rule.", Required: true},
									&cli.StringFlag{Name: "metadata", Usage: "The metadata.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/rules/%s/metadata/%s", cmd.String("project"), cmd.String("rule"), cmd.String("metadata"))
									fmt.Printf("Executing batch-delete on %s\n", parent)
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
