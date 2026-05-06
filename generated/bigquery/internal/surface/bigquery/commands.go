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

package bigquery

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
)

// Command returns the bigquery command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigquery",
		Usage: "manage BigQuery API resources",
		Commands: []*cli.Command{
			{
				Name:  "cancel",
				Usage: "Manage cancel resources",
				Commands: []*cli.Command{

					{
						Name:  "cancel-job",
						Usage: "cancel-job cancel",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The geographic location of the job.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("job_id"))
							fmt.Printf("Executing cancel-job on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "datasets",
				Usage: "Manage datasets resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe datasets",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "access-policy-version", Usage: "The version of the access policy schema to fetch.", Required: false},
							&cli.StringFlag{Name: "dataset-view", Usage: "Specifies the view that determines which dataset information is.", Required: false},
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "insert-dataset",
						Usage: "insert-dataset datasets",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "access-policy-version", Usage: "The version of the provided access policy schema.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing insert-dataset on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "patch-dataset",
						Usage: "patch-dataset datasets",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "access-policy-version", Usage: "The version of the provided access policy schema.", Required: false},
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "update-mode", Usage: "Specifies the fields of dataset that update/patch operation is.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing patch-dataset on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update datasets",
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "access-policy-version", Usage: "The version of the provided access policy schema.", Required: false},
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "update-mode", Usage: "Specifies the fields of dataset that update/patch operation is.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing update on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.BoolFlag{Name: "delete-contents", Usage: "If True, delete all the tables in the dataset.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing delete on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list datasets",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "all", Usage: "Whether to list all datasets, including hidden ones.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "An expression for filtering the results of the request by label.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "undelete",
						Usage: "undelete datasets",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing undelete on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "delete",
				Usage: "Manage delete resources",
				Commands: []*cli.Command{

					{
						Name:  "delete",
						Usage: "delete delete",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The geographic location of the job.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("job_id"))
							fmt.Printf("Executing delete on %s\n", project_id)
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
						Name:  "describe",
						Usage: "describe jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The geographic location of the job.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("job_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "insert-job",
						Usage: "insert-job jobs",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing insert-job on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list jobs",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "all-users", Usage: "Whether to display jobs owned by all users in the project.", Required: false},
							&cli.IntFlag{Name: "min-creation-time", Usage: "Min value for job creation time, in milliseconds since the POSIX epoch.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "parent-job-id", Usage: "If set, show only child jobs of the specified parent.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "projection", Usage: "Restrict information returned to a set of selected fields.", Required: false},
							&cli.StringSliceFlag{Name: "state-filter", Usage: "Filter for job state.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "models",
				Usage: "Manage models resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "model_id", Usage: "The ID of the model_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("model_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "patch-model",
						Usage: "patch-model models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "model_id", Usage: "The ID of the model_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("model_id"))
							fmt.Printf("Executing patch-model on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete models",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "model_id", Usage: "The ID of the model_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("model_id"))
							fmt.Printf("Executing delete on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "queries",
				Usage: "Manage queries resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "job_id", Usage: "The ID of the job_id.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The geographic location of the job.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("job_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "query",
						Usage: "query queries",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing query on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "routines",
				Usage: "Manage routines resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe routines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "routine_id", Usage: "The ID of the routine_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("routine_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "insert-routine",
						Usage: "insert-routine routines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing insert-routine on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update routines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "routine_id", Usage: "The ID of the routine_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("routine_id"))
							fmt.Printf("Executing update on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete routines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "routine_id", Usage: "The ID of the routine_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("routine_id"))
							fmt.Printf("Executing delete on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list routines",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "filter", Usage: "If set, then only the Routines matching this filter are returned.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "row-access-policies",
				Usage: "Manage row-access-policies resources",
				Commands: []*cli.Command{

					{
						Name:  "list",
						Usage: "list row-access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of results to return in a single response page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "describe",
						Usage: "describe row-access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "policy_id", Usage: "The ID of the policy_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"), cmd.String("policy_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "create",
						Usage: "create row-access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"))
							fmt.Printf("Executing create on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update row-access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "policy_id", Usage: "The ID of the policy_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"), cmd.String("policy_id"))
							fmt.Printf("Executing update on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete row-access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, it deletes the row access policy even if it's the last row.", Required: false},
							&cli.StringFlag{Name: "policy_id", Usage: "The ID of the policy_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"), cmd.String("policy_id"))
							fmt.Printf("Executing delete on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "batch-delete",
						Usage: "batch-delete row-access-policies",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "If set to true, it deletes the row access policy even if it's the last row.", Required: false},
							&cli.StringSliceFlag{Name: "policy-ids", Usage: "Policy IDs of the row access policies.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"))
							fmt.Printf("Executing batch-delete on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "service-account",
				Usage: "Manage service-account resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe service-account",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s", cmd.String("project_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},
				},
			},
			{
				Name:  "tables",
				Usage: "Manage tables resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "selected-fields", Usage: "List of table schema fields to return (comma-separated).", Required: false},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
							&cli.StringFlag{Name: "view", Usage: "Specifies the view that determines which table information is.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"))
							fmt.Printf("Executing describe on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "insert-table",
						Usage: "insert-table tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing insert-table on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "patch-table",
						Usage: "patch-table tables",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "autodetect-schema", Usage: "When true will autodetect schema, else will keep original schema.", Required: false},
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"))
							fmt.Printf("Executing patch-table on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update tables",
						Flags: []cli.Flag{
							&cli.BoolFlag{Name: "autodetect-schema", Usage: "When true will autodetect schema, else will keep original schema.", Required: false},
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"))
							fmt.Printf("Executing update on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "delete",
						Usage: "delete tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
							&cli.StringFlag{Name: "table_id", Usage: "The ID of the table_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s/%s", cmd.String("project_id"), cmd.String("dataset_id"), cmd.String("table_id"))
							fmt.Printf("Executing delete on %s\n", project_id)
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list tables",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "dataset_id", Usage: "The ID of the dataset_id.", Required: true},
							&cli.StringFlag{Name: "page-token", Usage: "Page token, returned by a previous call, to request the next page of.", Required: false},
							&cli.StringFlag{Name: "project_id", Usage: "The ID of the project_id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							project_id := fmt.Sprintf("%s/%s", cmd.String("project_id"), cmd.String("dataset_id"))
							fmt.Printf("Executing list on %s\n", project_id)
							return nil
						},
					},
				},
			},
		},
	}
}
