package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	configdelivery "cloud.google.com/go/configdelivery/apiv1"
	"cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
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
				Name:  "configdelivery",
				Usage: "manage Config Delivery API resources",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "project", Usage: "The project.", Required: false},
				},
				Commands: []*cli.Command{
					{
						Name:  "fleet-packages",
						Usage: "Manage fleet-packages resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list fleet-packages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configdeliverypb.ListFleetPackagesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListFleetPackages(ctx, req)
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
								Usage: "describe fleet-packages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.GetFleetPackageRequest{Name: name}
									resp, err := client.GetFleetPackage(ctx, req)
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
								Name:  "create",
								Usage: "create fleet-packages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet-package-id", Usage: "The fleet package id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.CreateFleetPackageRequest{Parent: parent}
									req.FleetPackageId = cmd.String("fleet-package-id")
									op, err := client.CreateFleetPackage(ctx, req)
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
								Usage: "update fleet-packages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.UpdateFleetPackageRequest{}
									req.FleetPackage = &configdeliverypb.FleetPackage{
										Name: name,
									}
									op, err := client.UpdateFleetPackage(ctx, req)
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
								Usage: "delete fleet-packages",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.DeleteFleetPackageRequest{Name: name}
									op, err := client.DeleteFleetPackage(ctx, req)
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
									client, err := configdelivery.NewClient(ctx)
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
									client, err := configdelivery.NewClient(ctx)
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
									client, err := configdelivery.NewClient(ctx)
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
									client, err := configdelivery.NewClient(ctx)
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
									client, err := configdelivery.NewClient(ctx)
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
									client, err := configdelivery.NewClient(ctx)
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
						Name:  "releases",
						Usage: "Manage releases resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configdeliverypb.ListReleasesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListReleases(ctx, req)
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
								Usage: "describe releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.GetReleaseRequest{Name: name}
									resp, err := client.GetRelease(ctx, req)
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
								Name:  "create",
								Usage: "create releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release-id", Usage: "The release id.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.CreateReleaseRequest{Parent: parent}
									req.ReleaseId = cmd.String("release-id")
									req.Release = &configdeliverypb.Release{
										Version: cmd.String("version"),
									}
									op, err := client.CreateRelease(ctx, req)
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
								Usage: "update releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "version", Usage: "The version.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.UpdateReleaseRequest{}
									req.Release = &configdeliverypb.Release{
										Name:    name,
										Version: cmd.String("version"),
									}
									var paths []string
									if cmd.IsSet("version") {
										paths = append(paths, "version")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateRelease(ctx, req)
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
								Usage: "delete releases",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.DeleteReleaseRequest{Name: name}
									op, err := client.DeleteRelease(ctx, req)
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
						},
					},
					{
						Name:  "resource-bundles",
						Usage: "Manage resource-bundles resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list resource-bundles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configdeliverypb.ListResourceBundlesRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListResourceBundles(ctx, req)
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
								Usage: "describe resource-bundles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.GetResourceBundleRequest{Name: name}
									resp, err := client.GetResourceBundle(ctx, req)
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
								Name:  "create",
								Usage: "create resource-bundles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource-bundle-id", Usage: "The resource bundle id.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.CreateResourceBundleRequest{Parent: parent}
									req.ResourceBundleId = cmd.String("resource-bundle-id")
									req.ResourceBundle = &configdeliverypb.ResourceBundle{
										Description: cmd.String("description"),
									}
									op, err := client.CreateResourceBundle(ctx, req)
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
								Usage: "update resource-bundles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "description", Usage: "The description.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.UpdateResourceBundleRequest{}
									req.ResourceBundle = &configdeliverypb.ResourceBundle{
										Name:        name,
										Description: cmd.String("description"),
									}
									var paths []string
									if cmd.IsSet("description") {
										paths = append(paths, "description")
									}
									req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
									op, err := client.UpdateResourceBundle(ctx, req)
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
								Usage: "delete resource-bundles",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.DeleteResourceBundleRequest{Name: name}
									op, err := client.DeleteResourceBundle(ctx, req)
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
						},
					},
					{
						Name:  "rollouts",
						Usage: "Manage rollouts resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configdeliverypb.ListRolloutsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListRollouts(ctx, req)
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
								Usage: "describe rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"), cmd.String("rollout"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.GetRolloutRequest{Name: name}
									resp, err := client.GetRollout(ctx, req)
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
								Name:  "suspend",
								Usage: "suspend rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
									&cli.StringFlag{Name: "reason", Usage: "The reason.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"), cmd.String("rollout"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.SuspendRolloutRequest{Name: name}
									req.Reason = cmd.String("reason")
									op, err := client.SuspendRollout(ctx, req)
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
								Name:  "resume",
								Usage: "resume rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
									&cli.StringFlag{Name: "reason", Usage: "The reason.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"), cmd.String("rollout"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.ResumeRolloutRequest{Name: name}
									req.Reason = cmd.String("reason")
									op, err := client.ResumeRollout(ctx, req)
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
								Name:  "abort",
								Usage: "abort rollouts",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "fleet_package", Usage: "The fleet_package.", Required: true},
									&cli.StringFlag{Name: "rollout", Usage: "The rollout.", Required: true},
									&cli.StringFlag{Name: "reason", Usage: "The reason.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/fleetPackages/%s/rollouts/%s", cmd.String("project"), cmd.String("location"), cmd.String("fleet_package"), cmd.String("rollout"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.AbortRolloutRequest{Name: name}
									req.Reason = cmd.String("reason")
									op, err := client.AbortRollout(ctx, req)
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
						Name:  "variants",
						Usage: "Manage variants resources",
						Commands: []*cli.Command{
							{
								Name:  "list",
								Usage: "list variants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
									&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
									&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
									&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									pageSize := cmd.Int("page-size")
									req := &configdeliverypb.ListVariantsRequest{Parent: parent}
									if pageSize > 0 {
										req.PageSize = int32(pageSize)
									}
									it := client.ListVariants(ctx, req)
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
								Usage: "describe variants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "variant", Usage: "The variant.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s/variants/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"), cmd.String("variant"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.GetVariantRequest{Name: name}
									resp, err := client.GetVariant(ctx, req)
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
								Name:  "create",
								Usage: "create variants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "variant-id", Usage: "The variant id.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									parent := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.CreateVariantRequest{Parent: parent}
									req.VariantId = cmd.String("variant-id")
									op, err := client.CreateVariant(ctx, req)
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
								Usage: "update variants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "variant", Usage: "The variant.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s/variants/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"), cmd.String("variant"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.UpdateVariantRequest{}
									req.Variant = &configdeliverypb.Variant{
										Name: name,
									}
									op, err := client.UpdateVariant(ctx, req)
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
								Usage: "delete variants",
								Flags: []cli.Flag{
									&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
									&cli.StringFlag{Name: "resource_bundle", Usage: "The resource_bundle.", Required: true},
									&cli.StringFlag{Name: "release", Usage: "The release.", Required: true},
									&cli.StringFlag{Name: "variant", Usage: "The variant.", Required: true},
								},
								Action: func(ctx context.Context, cmd *cli.Command) error {
									name := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s/variants/%s", cmd.String("project"), cmd.String("location"), cmd.String("resource_bundle"), cmd.String("release"), cmd.String("variant"))
									client, err := configdelivery.NewClient(ctx)
									if err != nil {
										return err
									}
									defer client.Close()
									req := &configdeliverypb.DeleteVariantRequest{Name: name}
									op, err := client.DeleteVariant(ctx, req)
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
