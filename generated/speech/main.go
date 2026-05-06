package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	speech "cloud.google.com/go/speech/apiv2"
	"cloud.google.com/go/speech/apiv2/speechpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func main() {
	app := &cli.Command{
		Name:  "gcloud",
		Usage: "Google Cloud CLI",
		Commands: []*cli.Command{
			{
				Name:  "speech",
				Usage: "manage Cloud Speech-to-Text API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "config",
						Usage: "Manage config resources",
						Commands: []*cli.Command{
							{
								Name:  "describe",
								Usage: "describe config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.GetConfigRequest{Name: name}
									resp, err := client.GetConfig(ctx, req)
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
								Usage: "update config",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "kms-key-name", Usage: "The kms key name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.UpdateConfigRequest{}
									req.Config = &speechpb.Config{
										Name:       name,
										KmsKeyName: cmd.String("kms-key-name"),
									}
									var paths []string
									if cmd.IsSet("kms-key-name") {
										paths = append(paths, "kms_key_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									resp, err := client.UpdateConfig(ctx, req)
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
						Name:  "custom-classes",
						Usage: "Manage custom-classes resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create custom-classes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "custom-class-id", Usage: "The custom class id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.CreateCustomClassRequest{Parent: parent}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.CustomClassId = cmd.String("custom-class-id")
									req.CustomClass = &speechpb.CustomClass{
										DisplayName: cmd.String("display-name"),
									}
									op, err := client.CreateCustomClass(ctx, req)
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
								Name:  "list",
								Usage: "list custom-classes",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe custom-classes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom_class", Usage: "The custom_class.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom_class"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.GetCustomClassRequest{Name: name}
									resp, err := client.GetCustomClass(ctx, req)
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
								Usage: "update custom-classes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom_class", Usage: "The custom_class.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom_class"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.UpdateCustomClassRequest{}
									req.CustomClass = &speechpb.CustomClass{
										Name:        name,
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateCustomClass(ctx, req)
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
								Name:  "delete",
								Usage: "delete custom-classes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom_class", Usage: "The custom_class.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom_class"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.DeleteCustomClassRequest{Name: name}
									op, err := client.DeleteCustomClass(ctx, req)
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
								Name:  "undelete",
								Usage: "undelete custom-classes",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "custom_class", Usage: "The custom_class.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom_class"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.UndeleteCustomClassRequest{Name: name}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.Etag = cmd.String("etag")
									op, err := client.UndeleteCustomClass(ctx, req)
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
						Name:  "locations",
						Usage: "Manage locations resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list locations",
								Flags: []cli.Flag{
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s", cmd.String("project"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &locationpb.ListLocationsRequest{Name: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListLocations(ctx, req)
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
								Usage: "describe locations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &locationpb.GetLocationRequest{Name: name}
									resp, err := client.GetLocation(ctx, req)
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
									client, err := speech.NewClient(ctx)
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
									client, err := speech.NewClient(ctx)
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
								Name:  "delete",
								Usage: "delete operations",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "operation", Usage: "The operation.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/operations/%s", cmd.String("project"), cmd.String("location"), cmd.String("operation"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &longrunningpb.DeleteOperationRequest{Name: name}
									if err := client.DeleteOperation(ctx, req); err != nil {
										return err
									}
									fmt.Printf("Deleted %s\n", name)
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
									client, err := speech.NewClient(ctx)
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
						Name:  "phrase-sets",
						Usage: "Manage phrase-sets resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create phrase-sets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "phrase-set-id", Usage: "The phrase set id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.CreatePhraseSetRequest{Parent: parent}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.PhraseSetId = cmd.String("phrase-set-id")
									req.PhraseSet = &speechpb.PhraseSet{
										DisplayName: cmd.String("display-name"),
									}
									op, err := client.CreatePhraseSet(ctx, req)
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
								Name:  "list",
								Usage: "list phrase-sets",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe phrase-sets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "phrase_set", Usage: "The phrase_set.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase_set"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.GetPhraseSetRequest{Name: name}
									resp, err := client.GetPhraseSet(ctx, req)
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
								Usage: "update phrase-sets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "phrase_set", Usage: "The phrase_set.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase_set"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.UpdatePhraseSetRequest{}
									req.PhraseSet = &speechpb.PhraseSet{
										Name:        name,
										DisplayName: cmd.String("display-name"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdatePhraseSet(ctx, req)
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
								Name:  "delete",
								Usage: "delete phrase-sets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "phrase_set", Usage: "The phrase_set.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase_set"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.DeletePhraseSetRequest{Name: name}
									op, err := client.DeletePhraseSet(ctx, req)
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
								Name:  "undelete",
								Usage: "undelete phrase-sets",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "phrase_set", Usage: "The phrase_set.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase_set"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.UndeletePhraseSetRequest{Name: name}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.Etag = cmd.String("etag")
									op, err := client.UndeletePhraseSet(ctx, req)
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
						Name:  "recognizers",
						Usage: "Manage recognizers resources",
						Commands: []*cli.Command{
							{
								Name:  "create",
								Usage: "create recognizers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "recognizer-id", Usage: "The recognizer id.", Required: false},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "model", Usage: "The model.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.CreateRecognizerRequest{Parent: parent}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.RecognizerId = cmd.String("recognizer-id")
									req.Recognizer = &speechpb.Recognizer{
										DisplayName: cmd.String("display-name"),
										Model:       cmd.String("model"),
									}
									op, err := client.CreateRecognizer(ctx, req)
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
								Name:  "list",
								Usage: "list recognizers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing list...")
									return nil
								},
							},
							{
								Name:  "describe",
								Usage: "describe recognizers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recognizer", Usage: "The recognizer.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.GetRecognizerRequest{Name: name}
									resp, err := client.GetRecognizer(ctx, req)
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
								Usage: "update recognizers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recognizer", Usage: "The recognizer.", Required: true},
									&cli.StringFlag{Name: "display-name", Usage: "The display name.", Required: false},
									&cli.StringFlag{Name: "model", Usage: "The model.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.UpdateRecognizerRequest{}
									req.Recognizer = &speechpb.Recognizer{
										Name:        name,
										DisplayName: cmd.String("display-name"),
										Model:       cmd.String("model"),
									}
									var paths []string
									if cmd.IsSet("display-name") {
										paths = append(paths, "display_name")
									}
									if cmd.IsSet("model") {
										paths = append(paths, "model")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateRecognizer(ctx, req)
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
								Name:  "delete",
								Usage: "delete recognizers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recognizer", Usage: "The recognizer.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.DeleteRecognizerRequest{Name: name}
									op, err := client.DeleteRecognizer(ctx, req)
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
								Name:  "undelete",
								Usage: "undelete recognizers",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "recognizer", Usage: "The recognizer.", Required: true},
									&cli.BoolFlag{Name: "validate-only", Usage: "The validate only.", Required: false},
									&cli.StringFlag{Name: "etag", Usage: "The etag.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
									client, err := speech.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &speechpb.UndeleteRecognizerRequest{Name: name}
									req.ValidateOnly = cmd.Bool("validate-only")
									req.Etag = cmd.String("etag")
									op, err := client.UndeleteRecognizer(ctx, req)
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
								Name:  "recognize",
								Usage: "recognize recognizers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing recognize...")
									return nil
								},
							},
							{
								Name:  "batch-recognize",
								Usage: "batch-recognize recognizers",
								Action: func(ctx context.Context, cmd *cli.Command) error {
									fmt.Println("Executing batch-recognize...")
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
