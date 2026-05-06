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

package policytroubleshooter

import (
	policytroubleshooter "cloud.google.com/go/policytroubleshooter/apiv1"
	"cloud.google.com/go/policytroubleshooter/apiv1/policytroubleshooterpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the policytroubleshooter command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "policytroubleshooter",
		Usage: "manage Policy Troubleshooter API resources",
		Commands: []*cli.Command{
			{
				Name:  "iam",
				Usage: "Manage iam resources",
				Commands: []*cli.Command{

					{
						Name:  "troubleshoot",
						Usage: "troubleshoot iam",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := policytroubleshooter.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &policytroubleshooterpb.TroubleshootIamPolicyRequest{}

							resp, err := client.TroubleshootIamPolicy(ctx, req)
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
