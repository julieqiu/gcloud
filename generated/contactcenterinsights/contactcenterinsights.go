package contactcenterinsights

import (
	"context"
	"fmt"
	"strings"

	contactcenterinsights "cloud.google.com/go/contactcenterinsights/apiv1"
	"cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud contactcenterinsights command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "contactcenterinsights",
		Usage: "manage Contact Center AI Insights API resources",
		Commands: []*cli.Command{
			{
				Name:  "analyses",
				Usage: "Manage analyses resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateAnalysisRequest{Parent: parent}
							req.Analysis = &contactcenterinsightspb.Analysis{
								Name: cmd.String("name"),
							}
							op, err := client.CreateAnalysis(ctx, req)
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
						Name:  "describe",
						Usage: "describe analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "analysis", Usage: "The analysis.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("analysis"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetAnalysisRequest{Name: name}
							resp, err := client.GetAnalysis(ctx, req)
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
						Usage: "list analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &contactcenterinsightspb.ListAnalysesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAnalyses(ctx, req)
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
						Name:  "delete",
						Usage: "delete analyses",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "analysis", Usage: "The analysis.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/analyses/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("analysis"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteAnalysisRequest{Name: name}
							if err := client.DeleteAnalysis(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "analysis-rules",
				Usage: "Manage analysis-rules resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "conversation-filter", Usage: "The conversation filter.", Required: false},
							&cli.BoolFlag{Name: "active", Usage: "The active.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateAnalysisRuleRequest{Parent: parent}
							req.AnalysisRule = &contactcenterinsightspb.AnalysisRule{
								DisplayName:        cmd.String("display-name"),
								ConversationFilter: cmd.String("conversation-filter"),
								Active:             cmd.Bool("active"),
							}
							resp, err := client.CreateAnalysisRule(ctx, req)
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
						Usage: "describe analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "analysis_rule", Usage: "The analysis_rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/analysisRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("analysis_rule"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetAnalysisRuleRequest{Name: name}
							resp, err := client.GetAnalysisRule(ctx, req)
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
						Usage: "list analysis-rules",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "analysis_rule", Usage: "The analysis_rule.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "conversation-filter", Usage: "The conversation filter.", Required: false},
							&cli.BoolFlag{Name: "active", Usage: "The active.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/analysisRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("analysis_rule"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateAnalysisRuleRequest{}
							req.AnalysisRule = &contactcenterinsightspb.AnalysisRule{
								Name:               name,
								DisplayName:        cmd.String("display-name"),
								ConversationFilter: cmd.String("conversation-filter"),
								Active:             cmd.Bool("active"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("conversation-filter") {
								paths = append(paths, "conversation_filter")
							}
							if cmd.IsSet("active") {
								paths = append(paths, "active")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateAnalysisRule(ctx, req)
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
						Usage: "delete analysis-rules",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "analysis_rule", Usage: "The analysis_rule.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/analysisRules/%s", cmd.String("project"), cmd.String("location"), cmd.String("analysis_rule"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteAnalysisRuleRequest{Name: name}
							if err := client.DeleteAnalysisRule(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "conversations",
				Usage: "Manage conversations resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation-id", Usage: "The conversation id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
							&cli.StringFlag{Name: "agent-id", Usage: "The agent id.", Required: false},
							&cli.StringFlag{Name: "metadata-json", Usage: "The metadata json.", Required: false},
							&cli.StringFlag{Name: "obfuscated-user-id", Usage: "The obfuscated user id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateConversationRequest{Parent: parent}
							req.ConversationId = cmd.String("conversation-id")
							req.Conversation = &contactcenterinsightspb.Conversation{
								Name:             cmd.String("name"),
								LanguageCode:     cmd.String("language-code"),
								AgentId:          cmd.String("agent-id"),
								MetadataJson:     cmd.String("metadata-json"),
								ObfuscatedUserId: cmd.String("obfuscated-user-id"),
							}
							resp, err := client.CreateConversation(ctx, req)
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
						Name:  "upload",
						Usage: "upload conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing upload on %s\n", parent)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
							&cli.StringFlag{Name: "agent-id", Usage: "The agent id.", Required: false},
							&cli.StringFlag{Name: "metadata-json", Usage: "The metadata json.", Required: false},
							&cli.StringFlag{Name: "obfuscated-user-id", Usage: "The obfuscated user id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateConversationRequest{}
							req.Conversation = &contactcenterinsightspb.Conversation{
								Name:             name,
								Name:             cmd.String("name"),
								LanguageCode:     cmd.String("language-code"),
								AgentId:          cmd.String("agent-id"),
								MetadataJson:     cmd.String("metadata-json"),
								ObfuscatedUserId: cmd.String("obfuscated-user-id"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("language-code") {
								paths = append(paths, "language_code")
							}
							if cmd.IsSet("agent-id") {
								paths = append(paths, "agent_id")
							}
							if cmd.IsSet("metadata-json") {
								paths = append(paths, "metadata_json")
							}
							if cmd.IsSet("obfuscated-user-id") {
								paths = append(paths, "obfuscated_user_id")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateConversation(ctx, req)
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
						Usage: "describe conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetConversationRequest{Name: name}
							resp, err := client.GetConversation(ctx, req)
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
						Usage: "list conversations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete conversations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteConversationRequest{Name: name}
							if err := client.DeleteConversation(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "bulk-analyze",
						Usage: "bulk-analyze conversations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing bulk-analyze...")
							return nil
						},
					},
					{
						Name:  "bulk-delete",
						Usage: "bulk-delete conversations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing bulk-delete...")
							return nil
						},
					},
					{
						Name:  "ingest",
						Usage: "ingest conversations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing ingest...")
							return nil
						},
					},
					{
						Name:  "calculate-stats",
						Usage: "calculate-stats conversations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing calculate-stats...")
							return nil
						},
					},
				},
			},
			{
				Name:  "encryption-spec",
				Usage: "Manage encryption-spec resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe encryption-spec",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/encryptionSpec", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetEncryptionSpecRequest{Name: name}
							resp, err := client.GetEncryptionSpec(ctx, req)
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
						Name:  "initialize",
						Usage: "initialize encryption-spec",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/encryptionSpec", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.InitializeEncryptionSpecRequest{Name: name}
							op, err := client.InitializeEncryptionSpec(ctx, req)
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
				Name:  "feedback-labels",
				Usage: "Manage feedback-labels resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "feedback-label-id", Usage: "The feedback label id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "labeled-resource", Usage: "The labeled resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/conversations/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateFeedbackLabelRequest{Parent: parent}
							req.FeedbackLabelId = cmd.String("feedback-label-id")
							req.FeedbackLabel = &contactcenterinsightspb.FeedbackLabel{
								Name:            cmd.String("name"),
								LabeledResource: cmd.String("labeled-resource"),
							}
							resp, err := client.CreateFeedbackLabel(ctx, req)
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
						Usage: "list feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &contactcenterinsightspb.ListFeedbackLabelsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListFeedbackLabels(ctx, req)
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
						Usage: "describe feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "feedback_label", Usage: "The feedback_label.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/feedbackLabels/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("feedback_label"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetFeedbackLabelRequest{Name: name}
							resp, err := client.GetFeedbackLabel(ctx, req)
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
						Usage: "update feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "feedback_label", Usage: "The feedback_label.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "labeled-resource", Usage: "The labeled resource.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/feedbackLabels/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("feedback_label"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateFeedbackLabelRequest{}
							req.FeedbackLabel = &contactcenterinsightspb.FeedbackLabel{
								Name:            name,
								Name:            cmd.String("name"),
								LabeledResource: cmd.String("labeled-resource"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("labeled-resource") {
								paths = append(paths, "labeled_resource")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateFeedbackLabel(ctx, req)
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
						Usage: "delete feedback-labels",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "conversation", Usage: "The conversation.", Required: true},
							&cli.StringFlag{Name: "feedback_label", Usage: "The feedback_label.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/conversations/%s/feedbackLabels/%s", cmd.String("project"), cmd.String("location"), cmd.String("conversation"), cmd.String("feedback_label"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteFeedbackLabelRequest{Name: name}
							if err := client.DeleteFeedbackLabel(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "insightsdata",
				Usage: "Manage insightsdata resources",
				Commands: []*cli.Command{
					{
						Name:  "export",
						Usage: "export insightsdata",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing export...")
							return nil
						},
					},
				},
			},
			{
				Name:  "issue-models",
				Usage: "Manage issue-models resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateIssueModelRequest{Parent: parent}
							req.IssueModel = &contactcenterinsightspb.IssueModel{
								Name:         cmd.String("name"),
								DisplayName:  cmd.String("display-name"),
								LanguageCode: cmd.String("language-code"),
							}
							op, err := client.CreateIssueModel(ctx, req)
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
						Name:  "update",
						Usage: "update issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateIssueModelRequest{}
							req.IssueModel = &contactcenterinsightspb.IssueModel{
								Name:         name,
								Name:         cmd.String("name"),
								DisplayName:  cmd.String("display-name"),
								LanguageCode: cmd.String("language-code"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("language-code") {
								paths = append(paths, "language_code")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateIssueModel(ctx, req)
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
						Usage: "describe issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetIssueModelRequest{Name: name}
							resp, err := client.GetIssueModel(ctx, req)
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
						Usage: "list issue-models",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteIssueModelRequest{Name: name}
							op, err := client.DeleteIssueModel(ctx, req)
							if err != nil {
								return err
							}
							if err := op.Wait(ctx); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "deploy",
						Usage: "deploy issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeployIssueModelRequest{Name: name}
							op, err := client.DeployIssueModel(ctx, req)
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
						Name:  "undeploy",
						Usage: "undeploy issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UndeployIssueModelRequest{Name: name}
							op, err := client.UndeployIssueModel(ctx, req)
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
						Name:  "export",
						Usage: "export issue-models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
							&cli.StringFlag{Name: "destination-object-uri", Usage: "The destination object uri.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.ExportIssueModelRequest{Name: name}
							req.Destination = &contactcenterinsightspb.ExportIssueModelRequest_GcsDestination{GcsDestination: &contactcenterinsightspb.GcsDestination{ObjectUri: cmd.String("destination-object-uri")}}
							op, err := client.ExportIssueModel(ctx, req)
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
						Name:  "import",
						Usage: "import issue-models",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing import...")
							return nil
						},
					},
					{
						Name:  "calculate-issue-model-stats",
						Usage: "calculate-issue-model-stats issue-models",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing calculate-issue-model-stats...")
							return nil
						},
					},
				},
			},
			{
				Name:  "issues",
				Usage: "Manage issues resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
							&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"), cmd.String("issue"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetIssueRequest{Name: name}
							resp, err := client.GetIssue(ctx, req)
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
						Usage: "list issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &contactcenterinsightspb.ListIssuesRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListIssues(ctx, req)
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
						Name:  "update",
						Usage: "update issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
							&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "display-description", Usage: "The display description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"), cmd.String("issue"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateIssueRequest{}
							req.Issue = &contactcenterinsightspb.Issue{
								Name:               name,
								Name:               cmd.String("name"),
								DisplayName:        cmd.String("display-name"),
								DisplayDescription: cmd.String("display-description"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("display-description") {
								paths = append(paths, "display_description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateIssue(ctx, req)
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
						Usage: "delete issues",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "issue_model", Usage: "The issue_model.", Required: true},
							&cli.StringFlag{Name: "issue", Usage: "The issue.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/issueModels/%s/issues/%s", cmd.String("project"), cmd.String("location"), cmd.String("issue_model"), cmd.String("issue"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteIssueRequest{Name: name}
							if err := client.DeleteIssue(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
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
						Name:  "query-metrics",
						Usage: "query-metrics locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing query-metrics...")
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "bulk-upload-feedback-labels",
						Usage: "bulk-upload-feedback-labels locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing bulk-upload-feedback-labels...")
							return nil
						},
					},
					{
						Name:  "bulk-download-feedback-labels",
						Usage: "bulk-download-feedback-labels locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing bulk-download-feedback-labels...")
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
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							client, err := contactcenterinsights.NewClient(ctx)
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
						Name:  "cancel",
						Usage: "cancel operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &longrunningpb.CancelOperationRequest{Name: name}
							if err := client.CancelOperation(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Cancelled %s\n", name)
							return nil
						},
					},
				},
			},
			{
				Name:  "phrase-matchers",
				Usage: "Manage phrase-matchers resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "version-tag", Usage: "The version tag.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.BoolFlag{Name: "active", Usage: "The active.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreatePhraseMatcherRequest{Parent: parent}
							req.PhraseMatcher = &contactcenterinsightspb.PhraseMatcher{
								Name:        cmd.String("name"),
								VersionTag:  cmd.String("version-tag"),
								DisplayName: cmd.String("display-name"),
								Active:      cmd.Bool("active"),
							}
							resp, err := client.CreatePhraseMatcher(ctx, req)
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
						Usage: "describe phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "phrase_matcher", Usage: "The phrase_matcher.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseMatchers/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase_matcher"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetPhraseMatcherRequest{Name: name}
							resp, err := client.GetPhraseMatcher(ctx, req)
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
						Usage: "list phrase-matchers",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "phrase_matcher", Usage: "The phrase_matcher.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseMatchers/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase_matcher"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeletePhraseMatcherRequest{Name: name}
							if err := client.DeletePhraseMatcher(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update phrase-matchers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "phrase_matcher", Usage: "The phrase_matcher.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "version-tag", Usage: "The version tag.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.BoolFlag{Name: "active", Usage: "The active.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseMatchers/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase_matcher"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdatePhraseMatcherRequest{}
							req.PhraseMatcher = &contactcenterinsightspb.PhraseMatcher{
								Name:        name,
								Name:        cmd.String("name"),
								VersionTag:  cmd.String("version-tag"),
								DisplayName: cmd.String("display-name"),
								Active:      cmd.Bool("active"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("version-tag") {
								paths = append(paths, "version_tag")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("active") {
								paths = append(paths, "active")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdatePhraseMatcher(ctx, req)
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
				Name:  "qa-questions",
				Usage: "Manage qa-questions resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
							&cli.StringFlag{Name: "qa-question-id", Usage: "The qa question id.", Required: false},
							&cli.StringFlag{Name: "abbreviation", Usage: "The abbreviation.", Required: false},
							&cli.StringFlag{Name: "question-body", Usage: "The question body.", Required: false},
							&cli.StringFlag{Name: "answer-instructions", Usage: "The answer instructions.", Required: false},
							&cli.IntFlag{Name: "order", Usage: "The order.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateQaQuestionRequest{Parent: parent}
							req.QaQuestionId = cmd.String("qa-question-id")
							req.QaQuestion = &contactcenterinsightspb.QaQuestion{
								Abbreviation:       cmd.String("abbreviation"),
								QuestionBody:       cmd.String("question-body"),
								AnswerInstructions: cmd.String("answer-instructions"),
								Order:              int32(cmd.Int("order")),
							}
							resp, err := client.CreateQaQuestion(ctx, req)
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
						Usage: "describe qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
							&cli.StringFlag{Name: "qa_question", Usage: "The qa_question.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s/qaQuestions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"), cmd.String("qa_question"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetQaQuestionRequest{Name: name}
							resp, err := client.GetQaQuestion(ctx, req)
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
						Usage: "update qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
							&cli.StringFlag{Name: "qa_question", Usage: "The qa_question.", Required: true},
							&cli.StringFlag{Name: "abbreviation", Usage: "The abbreviation.", Required: false},
							&cli.StringFlag{Name: "question-body", Usage: "The question body.", Required: false},
							&cli.StringFlag{Name: "answer-instructions", Usage: "The answer instructions.", Required: false},
							&cli.IntFlag{Name: "order", Usage: "The order.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s/qaQuestions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"), cmd.String("qa_question"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateQaQuestionRequest{}
							req.QaQuestion = &contactcenterinsightspb.QaQuestion{
								Name:               name,
								Abbreviation:       cmd.String("abbreviation"),
								QuestionBody:       cmd.String("question-body"),
								AnswerInstructions: cmd.String("answer-instructions"),
								Order:              int32(cmd.Int("order")),
							}
							var paths []string
							if cmd.IsSet("abbreviation") {
								paths = append(paths, "abbreviation")
							}
							if cmd.IsSet("question-body") {
								paths = append(paths, "question_body")
							}
							if cmd.IsSet("answer-instructions") {
								paths = append(paths, "answer_instructions")
							}
							if cmd.IsSet("order") {
								paths = append(paths, "order")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateQaQuestion(ctx, req)
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
						Usage: "delete qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
							&cli.StringFlag{Name: "qa_question", Usage: "The qa_question.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s/qaQuestions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"), cmd.String("qa_question"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteQaQuestionRequest{Name: name}
							if err := client.DeleteQaQuestion(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list qa-questions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &contactcenterinsightspb.ListQaQuestionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListQaQuestions(ctx, req)
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
				Name:  "qa-scorecards",
				Usage: "Manage qa-scorecards resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard-id", Usage: "The qa scorecard id.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateQaScorecardRequest{Parent: parent}
							req.QaScorecardId = cmd.String("qa-scorecard-id")
							req.QaScorecard = &contactcenterinsightspb.QaScorecard{
								DisplayName: cmd.String("display-name"),
								Description: cmd.String("description"),
							}
							resp, err := client.CreateQaScorecard(ctx, req)
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
						Usage: "describe qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetQaScorecardRequest{Name: name}
							resp, err := client.GetQaScorecard(ctx, req)
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
						Usage: "update qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateQaScorecardRequest{}
							req.QaScorecard = &contactcenterinsightspb.QaScorecard{
								Name:        name,
								DisplayName: cmd.String("display-name"),
								Description: cmd.String("description"),
							}
							var paths []string
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("description") {
								paths = append(paths, "description")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateQaScorecard(ctx, req)
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
						Usage: "delete qa-scorecards",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteQaScorecardRequest{Name: name}
							if err := client.DeleteQaScorecard(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list qa-scorecards",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
				},
			},
			{
				Name:  "revisions",
				Usage: "Manage revisions resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "qa-scorecard-revision-id", Usage: "The qa scorecard revision id.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateQaScorecardRevisionRequest{Parent: parent}
							req.QaScorecardRevisionId = cmd.String("qa-scorecard-revision-id")
							resp, err := client.CreateQaScorecardRevision(ctx, req)
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
						Usage: "describe revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetQaScorecardRevisionRequest{Name: name}
							resp, err := client.GetQaScorecardRevision(ctx, req)
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
						Name:  "tune-qa-scorecard-revision",
						Usage: "tune-qa-scorecard-revision revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
							&cli.StringFlag{Name: "parent", Usage: "The parent.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.TuneQaScorecardRevisionRequest{Name: name}
							req.Parent = cmd.String("parent")
							req.Filter = cmd.String("filter")
							req.ValidateOnly = cmd.Bool("validate-only")
							op, err := client.TuneQaScorecardRevision(ctx, req)
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
						Name:  "deploy",
						Usage: "deploy revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"))
							fmt.Printf("Executing deploy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "undeploy",
						Usage: "undeploy revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"))
							fmt.Printf("Executing undeploy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "qa_scorecard", Usage: "The qa_scorecard.", Required: true},
							&cli.StringFlag{Name: "revision", Usage: "The revision.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/qaScorecards/%s/revisions/%s", cmd.String("project"), cmd.String("location"), cmd.String("qa_scorecard"), cmd.String("revision"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteQaScorecardRevisionRequest{Name: name}
							if err := client.DeleteQaScorecardRevision(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list revisions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &contactcenterinsightspb.ListQaScorecardRevisionsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListQaScorecardRevisions(ctx, req)
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
				Name:  "settings",
				Usage: "Manage settings resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/settings", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetSettingsRequest{Name: name}
							resp, err := client.GetSettings(ctx, req)
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
						Usage: "update settings",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "language-code", Usage: "The language code.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/settings", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateSettingsRequest{}
							req.Settings = &contactcenterinsightspb.Settings{
								Name:         name,
								Name:         cmd.String("name"),
								LanguageCode: cmd.String("language-code"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("language-code") {
								paths = append(paths, "language_code")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateSettings(ctx, req)
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
				Name:  "views",
				Usage: "Manage views resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "value", Usage: "The value.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.CreateViewRequest{Parent: parent}
							req.View = &contactcenterinsightspb.View{
								Name:        cmd.String("name"),
								DisplayName: cmd.String("display-name"),
								Value:       cmd.String("value"),
							}
							resp, err := client.CreateView(ctx, req)
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
						Usage: "describe views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("view"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.GetViewRequest{Name: name}
							resp, err := client.GetView(ctx, req)
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
						Usage: "list views",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
							&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
							&cli.StringFlag{Name: "value", Usage: "The value.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("view"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.UpdateViewRequest{}
							req.View = &contactcenterinsightspb.View{
								Name:        name,
								Name:        cmd.String("name"),
								DisplayName: cmd.String("display-name"),
								Value:       cmd.String("value"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							if cmd.IsSet("display-name") {
								paths = append(paths, "display_name")
							}
							if cmd.IsSet("value") {
								paths = append(paths, "value")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateView(ctx, req)
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
						Usage: "delete views",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "The view.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/views/%s", cmd.String("project"), cmd.String("location"), cmd.String("view"))
							client, err := contactcenterinsights.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &contactcenterinsightspb.DeleteViewRequest{Name: name}
							if err := client.DeleteView(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
				},
			},
		},
	}
}
