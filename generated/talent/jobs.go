package jobs

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	talent "cloud.google.com/go/talent/apiv4"
	"cloud.google.com/go/talent/apiv4/talentpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud jobs command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "jobs",
		Usage: "manage Cloud Talent Solution API resources",
		Commands: []*cli.Command{
			{
				Name:  "client-events",
				Usage: "Manage client-events resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create client-events",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := talent.NewEventClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.CreateClientEventRequest{Parent: parent}
							resp, err := client.CreateClientEvent(ctx, req)
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
				Name:  "companies",
				Usage: "Manage companies resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: true},
							&cli.StringFlag{Name: "external-id", Usage: "The external id.", Required: true},
							&cli.StringFlag{Name: "headquarters-address", Usage: "The headquarters address.", Required: false},
							&cli.BoolFlag{Name: "hiring-agency", Usage: "The hiring agency.", Required: false},
							&cli.StringFlag{Name: "eeo-text", Usage: "The eeo text.", Required: false},
							&cli.StringFlag{Name: "website-uri", Usage: "The website uri.", Required: false},
							&cli.StringFlag{Name: "career-site-uri", Usage: "The career site uri.", Required: false},
							&cli.StringFlag{Name: "image-uri", Usage: "The image uri.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							client, err := talent.NewCompanyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.CreateCompanyRequest{Parent: parent}
							req.Company = &talentpb.Company{
								Name:                cmd.String("name"),
								DisplayName:         cmd.String("display-name"),
								ExternalId:          cmd.String("external-id"),
								HeadquartersAddress: cmd.String("headquarters-address"),
								HiringAgency:        cmd.Bool("hiring-agency"),
								EeoText:             cmd.String("eeo-text"),
								WebsiteUri:          cmd.String("website-uri"),
								CareerSiteUri:       cmd.String("career-site-uri"),
								ImageUri:            cmd.String("image-uri"),
							}
							resp, err := client.CreateCompany(ctx, req)
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
						Usage: "describe companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "company", Usage: "The company.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/companies/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("company"))
							client, err := talent.NewCompanyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.GetCompanyRequest{Name: name}
							resp, err := client.GetCompany(ctx, req)
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
						Usage: "update companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "company", Usage: "The company.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "external-id", Usage: "The external id.", Required: false},
							&cli.StringFlag{Name: "headquarters-address", Usage: "The headquarters address.", Required: false},
							&cli.BoolFlag{Name: "hiring-agency", Usage: "The hiring agency.", Required: false},
							&cli.StringFlag{Name: "eeo-text", Usage: "The eeo text.", Required: false},
							&cli.StringFlag{Name: "website-uri", Usage: "The website uri.", Required: false},
							&cli.StringFlag{Name: "career-site-uri", Usage: "The career site uri.", Required: false},
							&cli.StringFlag{Name: "image-uri", Usage: "The image uri.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/companies/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("company"))
							client, err := talent.NewCompanyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.UpdateCompanyRequest{}
							req.Company = &talentpb.Company{
								Name:                name,
								Name:                cmd.String("name"),
								DisplayName:         cmd.String("display-name"),
								ExternalId:          cmd.String("external-id"),
								HeadquartersAddress: cmd.String("headquarters-address"),
								HiringAgency:        cmd.Bool("hiring-agency"),
								EeoText:             cmd.String("eeo-text"),
								WebsiteUri:          cmd.String("website-uri"),
								CareerSiteUri:       cmd.String("career-site-uri"),
								ImageUri:            cmd.String("image-uri"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("external-id") {
								paths = append(paths, "external_id")
							}
							if cmd.IsSet("headquarters-address") {
								paths = append(paths, "headquarters_address")
							}
							if cmd.IsSet("hiring-agency") {
								paths = append(paths, "hiring_agency")
							}
							if cmd.IsSet("eeo-text") {
								paths = append(paths, "eeo_text")
							}
							if cmd.IsSet("website-uri") {
								paths = append(paths, "website_uri")
							}
							if cmd.IsSet("career-site-uri") {
								paths = append(paths, "career_site_uri")
							}
							if cmd.IsSet("image-uri") {
								paths = append(paths, "image_uri")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateCompany(ctx, req)
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
						Usage: "delete companies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "company", Usage: "The company.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/companies/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("company"))
							client, err := talent.NewCompanyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.DeleteCompanyRequest{Name: name}
							if err := client.DeleteCompany(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list companies",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := talent.NewCompanyClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &talentpb.ListCompaniesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListCompanies(ctx, req)
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
				Name:  "jobs",
				Usage: "Manage jobs resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "company", Usage: "The company.", Required: true},
							&cli.StringFlag{Name: "requisition-id", Usage: "The requisition id.", Required: true},
							&cli.StringFlag{Name: "title", Usage: "The title.", Required: true},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: true},
							&cli.StringFlag{Name: "department", Usage: "The department.", Required: false},
							&cli.StringFlag{Name: "incentives", Usage: "The incentives.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
							&cli.IntFlag{Name: "promotion-value", Usage: "The promotion value.", Required: false},
							&cli.StringFlag{Name: "qualifications", Usage: "The qualifications.", Required: false},
							&cli.StringFlag{Name: "responsibilities", Usage: "The responsibilities.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							client, err := talent.NewJobClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.CreateJobRequest{Parent: parent}
							req.Job = &talentpb.Job{
								Name:             cmd.String("name"),
								Company:          cmd.String("company"),
								RequisitionId:    cmd.String("requisition-id"),
								Title:            cmd.String("title"),
								Description:      cmd.String("description"),
								Department:       cmd.String("department"),
								Incentives:       cmd.String("incentives"),
								LanguageCode:     cmd.String("language-code"),
								PromotionValue:   int32(cmd.Int("promotion-value")),
								Qualifications:   cmd.String("qualifications"),
								Responsibilities: cmd.String("responsibilities"),
							}
							resp, err := client.CreateJob(ctx, req)
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
						Name:  "batch-create",
						Usage: "batch-create jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							fmt.Printf("Executing batch-create on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/jobs/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("job"))
							client, err := talent.NewJobClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.GetJobRequest{Name: name}
							resp, err := client.GetJob(ctx, req)
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
						Usage: "update jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "company", Usage: "The company.", Required: false},
							&cli.StringFlag{Name: "requisition-id", Usage: "The requisition id.", Required: false},
							&cli.StringFlag{Name: "title", Usage: "The title.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
							&cli.StringFlag{Name: "department", Usage: "The department.", Required: false},
							&cli.StringFlag{Name: "incentives", Usage: "The incentives.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
							&cli.IntFlag{Name: "promotion-value", Usage: "The promotion value.", Required: false},
							&cli.StringFlag{Name: "qualifications", Usage: "The qualifications.", Required: false},
							&cli.StringFlag{Name: "responsibilities", Usage: "The responsibilities.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/jobs/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("job"))
							client, err := talent.NewJobClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.UpdateJobRequest{}
							req.Job = &talentpb.Job{
								Name:             name,
								Name:             cmd.String("name"),
								Company:          cmd.String("company"),
								RequisitionId:    cmd.String("requisition-id"),
								Title:            cmd.String("title"),
								Description:      cmd.String("description"),
								Department:       cmd.String("department"),
								Incentives:       cmd.String("incentives"),
								LanguageCode:     cmd.String("language-code"),
								PromotionValue:   int32(cmd.Int("promotion-value")),
								Qualifications:   cmd.String("qualifications"),
								Responsibilities: cmd.String("responsibilities"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("company") {
								paths = append(paths, "company")
							}
							if cmd.IsSet("requisition-id") {
								paths = append(paths, "requisition_id")
							}
							if cmd.IsSet("title") {
								paths = append(paths, "title")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							if cmd.IsSet("department") {
								paths = append(paths, "department")
							}
							if cmd.IsSet("incentives") {
								paths = append(paths, "incentives")
							}
							if cmd.IsSet("language-code") {
								paths = append(paths, "language_code")
							}
							if cmd.IsSet("promotion-value") {
								paths = append(paths, "promotion_value")
							}
							if cmd.IsSet("qualifications") {
								paths = append(paths, "qualifications")
							}
							if cmd.IsSet("responsibilities") {
								paths = append(paths, "responsibilities")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateJob(ctx, req)
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
						Name:  "batch-update",
						Usage: "batch-update jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							fmt.Printf("Executing batch-update on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "job", Usage: "The job.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s/jobs/%s", cmd.String("project"), cmd.String("tenant"), cmd.String("job"))
							client, err := talent.NewJobClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.DeleteJobRequest{Name: name}
							if err := client.DeleteJob(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "batch-delete",
						Usage: "batch-delete jobs",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing batch-delete on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							client, err := talent.NewJobClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &talentpb.ListJobsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListJobs(ctx, req)
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
						Name:  "search",
						Usage: "search jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							fmt.Printf("Executing search on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "search-for-alert",
						Usage: "search-for-alert jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							fmt.Printf("Executing search-for-alert on %s\n", parent)
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
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := talent.NewCompanyClient(ctx)
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := talent.NewCompletionClient(ctx)
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := talent.NewEventClient(ctx)
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := talent.NewJobClient(ctx)
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
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/operations/%s", cmd.String("project"), cmd.String("operation"))
							client, err := talent.NewTenantClient(ctx)
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
				Name:  "tenants",
				Usage: "Manage tenants resources",
				Commands: []*cli.Command{
					{
						Name:  "complete-query",
						Usage: "complete-query tenants",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing complete-query...")
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "create tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "external-id", Usage: "The external id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := talent.NewTenantClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.CreateTenantRequest{Parent: parent}
							req.Tenant = &talentpb.Tenant{
								Name:       cmd.String("name"),
								ExternalId: cmd.String("external-id"),
							}
							resp, err := client.CreateTenant(ctx, req)
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
						Usage: "describe tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							client, err := talent.NewTenantClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.GetTenantRequest{Name: name}
							resp, err := client.GetTenant(ctx, req)
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
						Usage: "update tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "external-id", Usage: "The external id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							client, err := talent.NewTenantClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.UpdateTenantRequest{}
							req.Tenant = &talentpb.Tenant{
								Name:       name,
								Name:       cmd.String("name"),
								ExternalId: cmd.String("external-id"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("external-id") {
								paths = append(paths, "external_id")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateTenant(ctx, req)
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
						Usage: "delete tenants",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "tenant", Usage: "The tenant.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/tenants/%s", cmd.String("project"), cmd.String("tenant"))
							client, err := talent.NewTenantClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &talentpb.DeleteTenantRequest{Name: name}
							if err := client.DeleteTenant(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list tenants",
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
