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

package speech

import (
	speech "cloud.google.com/go/speech/apiv2"
	"cloud.google.com/go/speech/apiv2/speechpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the speech command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "speech",
		Usage: "manage Cloud Speech-to-Text API resources",
		Commands: []*cli.Command{
			{
				Name:  "config",
				Usage: "Manage config resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.GetConfigRequest{
								Name: name,
							}

							resp, err := client.GetConfig(ctx, req)
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
						Usage: "update config",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "config.name" not yet supported.
							config_name := fmt.Sprintf("projects/%s/locations/%s/config", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", config_name)
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
							&cli.StringFlag{Name: "custom-class-id", Usage: "The ID to use for the CustomClass, which will become the final component of.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the CustomClass, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.CreateCustomClassRequest{
								Parent:        parent,
								ValidateOnly:  cmd.Bool("validate-only"),
								CustomClassId: cmd.String("custom-class-id"),
							}

							op, err := client.CreateCustomClass(ctx, req)
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
						Name:  "list",
						Usage: "list custom-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "Number of results per requests.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "Whether, or not, to show resources that have been deleted.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.ListCustomClassesRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListCustomClasses(ctx, req)
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
						Name:  "describe",
						Usage: "describe custom-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-classe", Usage: "The ID of the custom classe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-classe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.GetCustomClassRequest{
								Name: name,
							}

							resp, err := client.GetCustomClass(ctx, req)
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
						Usage: "update custom-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-classe", Usage: "The ID of the custom classe.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the updated CustomClass, but do.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "custom_class.name" not yet supported.
							custom_class_name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-classe"))
							fmt.Printf("Executing update on %s\n", custom_class_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete custom-classes",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the CustomClass is not found, the request will succeed.", Required: false},
							&cli.StringFlag{Name: "custom-classe", Usage: "The ID of the custom classe.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of other.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the deleted CustomClass, but do.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-classe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.DeleteCustomClassRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								AllowMissing: cmd.Bool("allow-missing"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteCustomClass(ctx, req)
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
						Name:  "undelete",
						Usage: "undelete custom-classes",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "custom-classe", Usage: "The ID of the custom classe.", Required: true},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of other.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the undeleted CustomClass, but do.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/customClasses/%s", cmd.String("project"), cmd.String("location"), cmd.String("custom-classe"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.UndeleteCustomClassRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.UndeleteCustomClass(ctx, req)
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
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s", cmd.String("project"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
				Name:  "operations",
				Usage: "Manage operations resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "filter", Usage: "The filter.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The page size.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The page token.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "return-partial-success", Usage: "The return partial success.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing list on %s\n", name)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe operations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "operation", Usage: "The ID of the operation.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
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
				Name:  "phrase-sets",
				Usage: "Manage phrase-sets resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create phrase-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-set-id", Usage: "The ID to use for the PhraseSet, which will become the final component of.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the PhraseSet, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.CreatePhraseSetRequest{
								Parent:       parent,
								ValidateOnly: cmd.Bool("validate-only"),
								PhraseSetId:  cmd.String("phrase-set-id"),
							}

							op, err := client.CreatePhraseSet(ctx, req)
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
						Name:  "list",
						Usage: "list phrase-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of PhraseSets to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "Whether, or not, to show resources that have been deleted.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.ListPhraseSetsRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListPhraseSets(ctx, req)
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
						Name:  "describe",
						Usage: "describe phrase-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-set", Usage: "The ID of the phrase set.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase-set"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.GetPhraseSetRequest{
								Name: name,
							}

							resp, err := client.GetPhraseSet(ctx, req)
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
						Usage: "update phrase-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-set", Usage: "The ID of the phrase set.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the updated PhraseSet, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "phrase_set.name" not yet supported.
							phrase_set_name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase-set"))
							fmt.Printf("Executing update on %s\n", phrase_set_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete phrase-sets",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the PhraseSet is not found, the request will succeed.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of other.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-set", Usage: "The ID of the phrase set.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the deleted PhraseSet, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase-set"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.DeletePhraseSetRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								AllowMissing: cmd.Bool("allow-missing"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeletePhraseSet(ctx, req)
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
						Name:  "undelete",
						Usage: "undelete phrase-sets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of other.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "phrase-set", Usage: "The ID of the phrase set.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the undeleted PhraseSet, but do.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/phraseSets/%s", cmd.String("project"), cmd.String("location"), cmd.String("phrase-set"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.UndeletePhraseSetRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.UndeletePhraseSet(ctx, req)
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
				Name:  "recognizers",
				Usage: "Manage recognizers resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create recognizers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recognizer-id", Usage: "The ID to use for the Recognizer, which will become the final component of.", Required: false},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the Recognizer, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.CreateRecognizerRequest{
								Parent:       parent,
								ValidateOnly: cmd.Bool("validate-only"),
								RecognizerId: cmd.String("recognizer-id"),
							}

							op, err := client.CreateRecognizer(ctx, req)
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
						Name:  "list",
						Usage: "list recognizers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of Recognizers to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "A page token, received from a previous.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.BoolFlag{Name: "show-deleted", Usage: "Whether, or not, to show resources that have been deleted.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.ListRecognizersRequest{
								Parent:      parent,
								PageSize:    int32(cmd.Int("page-size")),
								PageToken:   cmd.String("page-token"),
								ShowDeleted: cmd.Bool("show-deleted"),
							}

							limit := cmd.Int("limit")
							it := client.ListRecognizers(ctx, req)
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
						Name:  "describe",
						Usage: "describe recognizers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recognizer", Usage: "The ID of the recognizer.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.GetRecognizerRequest{
								Name: name,
							}

							resp, err := client.GetRecognizer(ctx, req)
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
						Usage: "update recognizers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recognizer", Usage: "The ID of the recognizer.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the updated Recognizer, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "recognizer.name" not yet supported.
							recognizer_name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
							fmt.Printf("Executing update on %s\n", recognizer_name)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete recognizers",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-missing", Usage: "If set to true, and the Recognizer is not found, the request will succeed.", Required: false},
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of other.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recognizer", Usage: "The ID of the recognizer.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the deleted Recognizer, but do not.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.DeleteRecognizerRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								AllowMissing: cmd.Bool("allow-missing"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.DeleteRecognizer(ctx, req)
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
						Name:  "undelete",
						Usage: "undelete recognizers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "etag", Usage: "This checksum is computed by the server based on the value of other.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recognizer", Usage: "The ID of the recognizer.", Required: true},
							&cli.BoolFlag{Name: "validate-only", Usage: "If set, validate the request and preview the undeleted Recognizer, but do.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := speech.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &speechpb.UndeleteRecognizerRequest{
								Name:         name,
								ValidateOnly: cmd.Bool("validate-only"),
								Etag:         cmd.String("etag"),
							}

							op, err := client.UndeleteRecognizer(ctx, req)
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
						Name:  "recognize",
						Usage: "recognize recognizers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recognizer", Usage: "The ID of the recognizer.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							recognizer := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
							fmt.Printf("Executing recognize on %s\n", recognizer)
							return nil
						},
					},

					{
						Name:  "batch-recognize",
						Usage: "batch-recognize recognizers",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "processing-strategy", Usage: "Processing strategy to use for this request.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "recognizer", Usage: "The ID of the recognizer.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							recognizer := fmt.Sprintf("projects/%s/locations/%s/recognizers/%s", cmd.String("project"), cmd.String("location"), cmd.String("recognizer"))
							fmt.Printf("Executing batch-recognize on %s\n", recognizer)
							return nil
						},
					},
				},
			},
		},
	}
}
