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

package language

import (
	language "cloud.google.com/go/language/apiv2"
	"cloud.google.com/go/language/apiv2/languagepb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the language command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "language",
		Usage: "manage Cloud Natural Language API resources",
		Commands: []*cli.Command{
			{
				Name:  "documents",
				Usage: "Manage documents resources",
				Commands: []*cli.Command{

					{
						Name:  "analyze-sentiment",
						Usage: "analyze-sentiment documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encoding-type", Usage: "The encoding type used by the API to calculate sentence offsets.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := language.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &languagepb.AnalyzeSentimentRequest{
								EncodingType: languagepb.EncodingType(languagepb.EncodingType_value[cmd.String("encoding-type")]),
							}

							resp, err := client.AnalyzeSentiment(ctx, req)
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
						Name:  "analyze-entities",
						Usage: "analyze-entities documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encoding-type", Usage: "The encoding type used by the API to calculate offsets.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := language.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &languagepb.AnalyzeEntitiesRequest{
								EncodingType: languagepb.EncodingType(languagepb.EncodingType_value[cmd.String("encoding-type")]),
							}

							resp, err := client.AnalyzeEntities(ctx, req)
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
						Name:  "classify-text",
						Usage: "classify-text documents",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := language.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &languagepb.ClassifyTextRequest{}

							resp, err := client.ClassifyText(ctx, req)
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
						Name:  "moderate-text",
						Usage: "moderate-text documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "model-version", Usage: "The model version to use for ModerateText.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := language.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &languagepb.ModerateTextRequest{
								ModelVersion: languagepb.ModerateTextRequest_ModelVersion(languagepb.ModerateTextRequest_ModelVersion_value[cmd.String("model-version")]),
							}

							resp, err := client.ModerateText(ctx, req)
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
						Name:  "annotate-text",
						Usage: "annotate-text documents",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "encoding-type", Usage: "The encoding type used by the API to calculate offsets.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := language.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &languagepb.AnnotateTextRequest{
								EncodingType: languagepb.EncodingType(languagepb.EncodingType_value[cmd.String("encoding-type")]),
							}

							resp, err := client.AnnotateText(ctx, req)
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
		},
	}
}
