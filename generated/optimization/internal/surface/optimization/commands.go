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

package optimization

import (
	optimization "cloud.google.com/go/optimization/apiv1"
	"cloud.google.com/go/optimization/apiv1/optimizationpb"
	"context"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
)

// Command returns the cloudoptimization command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "cloudoptimization",
		Usage: "manage Cloud Optimization API resources",
		Commands: []*cli.Command{
			{
				Name:  "locations",
				Usage: "Manage locations resources",
				Commands: []*cli.Command{

					{
						Name:  "optimize-tours",
						Usage: "optimize-tours locations",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "allow-large-deadline-despite-interruption-risk", Usage: "If this is set, then the request can have a deadline.", Required: false},
							&cli.BoolFlag{Name: "consider-road-traffic", Usage: "Consider traffic estimation in calculating `ShipmentRoute` fields.", Required: false},
							&cli.FloatFlag{Name: "geodesic-meters-per-second", Usage: "When `use_geodesic_distances` is true, this field must be set and defines.", Required: false},
							&cli.BoolFlag{Name: "interpret-injected-solutions-using-labels", Usage: "If true:.", Required: false},
							&cli.StringFlag{Name: "label", Usage: "Label that may be used to identify this request, reported back in the.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "max-validation-errors", Usage: "Truncates the number of validation errors returned.", Required: false},
							&cli.BoolFlag{Name: "populate-polylines", Usage: "If true, polylines will be populated in response `ShipmentRoute`s.", Required: false},
							&cli.BoolFlag{Name: "populate-transition-polylines", Usage: "If true, polylines will be populated in response.", Required: false},
							&cli.BoolFlag{Name: "populate-travel-step-polylines", Usage: "Deprecated: Use.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "search-mode", Usage: "Search mode used to solve the request.", Required: false},
							&cli.StringFlag{Name: "solving-mode", Usage: "By default, the solving mode is `DEFAULT_SOLVE` (0).", Required: false},
							&cli.BoolFlag{Name: "use-geodesic-distances", Usage: "If true, travel distances will be computed using geodesic distances instead.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := optimization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &optimizationpb.OptimizeToursRequest{
								Parent:                                parent,
								SolvingMode:                           optimizationpb.OptimizeToursRequest_SolvingMode(optimizationpb.OptimizeToursRequest_SolvingMode_value[cmd.String("solving-mode")]),
								SearchMode:                            optimizationpb.OptimizeToursRequest_SearchMode(optimizationpb.OptimizeToursRequest_SearchMode_value[cmd.String("search-mode")]),
								InterpretInjectedSolutionsUsingLabels: cmd.Bool("interpret-injected-solutions-using-labels"),
								ConsiderRoadTraffic:                   cmd.Bool("consider-road-traffic"),
								PopulatePolylines:                     cmd.Bool("populate-polylines"),
								PopulateTransitionPolylines:           cmd.Bool("populate-transition-polylines"),
								AllowLargeDeadlineDespiteInterruptionRisk: cmd.Bool("allow-large-deadline-despite-interruption-risk"),
								UseGeodesicDistances:                      cmd.Bool("use-geodesic-distances"),
								GeodesicMetersPerSecond:                   runtime.Ptr(cmd.Float("geodesic-meters-per-second")),
								MaxValidationErrors:                       runtime.Ptr(int32(cmd.Int("max-validation-errors"))),
								Label:                                     cmd.String("label"),
								PopulateTravelStepPolylines:               cmd.Bool("populate-travel-step-polylines"),
							}

							resp, err := client.OptimizeTours(ctx, req)
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
						Name:  "batch-optimize-tours",
						Usage: "batch-optimize-tours locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := optimization.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &optimizationpb.BatchOptimizeToursRequest{
								Parent: parent,
							}

							op, err := client.BatchOptimizeTours(ctx, req)
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
		},
	}
}
