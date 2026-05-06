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

package reservation

import (
	reservation "cloud.google.com/go/reservation/apiv1"
	"cloud.google.com/go/reservation/apiv1/reservationpb"
	"context"
	"errors"
	"fmt"
	"gcloud/internal/auth"
	"gcloud/internal/runtime"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
)

// Command returns the bigqueryreservation command tree for inclusion under the gcloud root.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "bigqueryreservation",
		Usage: "manage BigQuery Reservation API resources",
		Commands: []*cli.Command{
			{
				Name:  "assignments",
				Usage: "Manage assignments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "assignment-id", Usage: "The optional assignment ID.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.CreateAssignmentRequest{
								Parent:       parent,
								AssignmentId: cmd.String("assignment-id"),
							}

							resp, err := client.CreateAssignment(ctx, req)
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
						Usage: "list assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.ListAssignmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListAssignments(ctx, req)
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
						Name:  "delete",
						Usage: "delete assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "assignment", Usage: "The ID of the assignment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s/assignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"), cmd.String("assignment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteAssignment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.DeleteAssignmentRequest{
								Name: name,
							}

							if err := client.DeleteAssignment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "move",
						Usage: "move assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "assignment", Usage: "The ID of the assignment.", Required: true},
							&cli.StringFlag{Name: "destination-id", Usage: "The new reservation ID, e.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s/assignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"), cmd.String("assignment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.MoveAssignmentRequest{
								Name:          name,
								DestinationId: cmd.String("destination-id"),
							}

							resp, err := client.MoveAssignment(ctx, req)
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
						Usage: "update assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "assignment", Usage: "The ID of the assignment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "assignment.name" not yet supported.
							assignment_name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s/assignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"), cmd.String("assignment"))
							fmt.Printf("Executing update on %s\n", assignment_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "bi-reservation",
				Usage: "Manage bi-reservation resources",
				Commands: []*cli.Command{

					{
						Name:  "describe",
						Usage: "describe bi-reservation",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/biReservation", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.GetBiReservationRequest{
								Name: name,
							}

							resp, err := client.GetBiReservation(ctx, req)
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
						Usage: "update bi-reservation",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "bi_reservation.name" not yet supported.
							bi_reservation_name := fmt.Sprintf("projects/%s/locations/%s/biReservation", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing update on %s\n", bi_reservation_name)
							return nil
						},
					},
				},
			},
			{
				Name:  "capacity-commitments",
				Usage: "Manage capacity-commitments resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "capacity-commitment-id", Usage: "The optional capacity commitment ID.", Required: false},
							&cli.BoolFlag{Name: "enforce-single-admin-project-per-org", Usage: "If true, fail the request if another project in the organization has a.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.CreateCapacityCommitmentRequest{
								Parent:                          parent,
								EnforceSingleAdminProjectPerOrg: cmd.Bool("enforce-single-admin-project-per-org"),
								CapacityCommitmentId:            cmd.String("capacity-commitment-id"),
							}

							resp, err := client.CreateCapacityCommitment(ctx, req)
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
						Usage: "list capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.ListCapacityCommitmentsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListCapacityCommitments(ctx, req)
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
						Usage: "describe capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "capacity-commitment", Usage: "The ID of the capacity commitment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity-commitment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.GetCapacityCommitmentRequest{
								Name: name,
							}

							resp, err := client.GetCapacityCommitment(ctx, req)
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
						Name:  "delete",
						Usage: "delete capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "capacity-commitment", Usage: "The ID of the capacity commitment.", Required: true},
							&cli.BoolFlag{Name: "force", Usage: "Can be used to force delete commitments even if assignments exist.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity-commitment"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteCapacityCommitment on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.DeleteCapacityCommitmentRequest{
								Name:  name,
								Force: cmd.Bool("force"),
							}

							if err := client.DeleteCapacityCommitment(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "capacity-commitment", Usage: "The ID of the capacity commitment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "capacity_commitment.name" not yet supported.
							capacity_commitment_name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity-commitment"))
							fmt.Printf("Executing update on %s\n", capacity_commitment_name)
							return nil
						},
					},

					{
						Name:  "split",
						Usage: "split capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "capacity-commitment", Usage: "The ID of the capacity commitment.", Required: true},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.IntFlag{Name: "slot-count", Usage: "Number of slots in the capacity commitment after the split.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity-commitment"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.SplitCapacityCommitmentRequest{
								Name:      name,
								SlotCount: cmd.Int("slot-count"),
							}

							resp, err := client.SplitCapacityCommitment(ctx, req)
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
						Name:  "merge",
						Usage: "merge capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "capacity-commitment-id", Usage: "The optional resulting capacity commitment ID.", Required: false},
							&cli.StringSliceFlag{Name: "capacity-commitment-ids", Usage: "Ids of capacity commitments to merge.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.MergeCapacityCommitmentsRequest{
								Parent:                parent,
								CapacityCommitmentIds: cmd.StringSlice("capacity-commitment-ids"),
								CapacityCommitmentId:  cmd.String("capacity-commitment-id"),
							}

							resp, err := client.MergeCapacityCommitments(ctx, req)
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
						Name:  "search-assignments",
						Usage: "search-assignments locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Please specify resource name as assignee in the query.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.SearchAssignmentsRequest{
								Parent:    parent,
								Query:     cmd.String("query"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchAssignments(ctx, req)
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
						Name:  "search-all-assignments",
						Usage: "search-all-assignments locations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "query", Usage: "Please specify resource name as assignee in the query.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.SearchAllAssignmentsRequest{
								Parent:    parent,
								Query:     cmd.String("query"),
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.SearchAllAssignments(ctx, req)
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
				},
			},
			{
				Name:  "reservation-groups",
				Usage: "Manage reservation-groups resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create reservation-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation-group-id", Usage: "The reservation group ID.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.CreateReservationGroupRequest{
								Parent:             parent,
								ReservationGroupId: cmd.String("reservation-group-id"),
							}

							resp, err := client.CreateReservationGroup(ctx, req)
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
						Name:  "describe",
						Usage: "describe reservation-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation-group", Usage: "The ID of the reservation group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservationGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation-group"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.GetReservationGroupRequest{
								Name: name,
							}

							resp, err := client.GetReservationGroup(ctx, req)
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
						Name:  "delete",
						Usage: "delete reservation-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation-group", Usage: "The ID of the reservation group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservationGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation-group"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteReservationGroup on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.DeleteReservationGroupRequest{
								Name: name,
							}

							if err := client.DeleteReservationGroup(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "list",
						Usage: "list reservation-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.ListReservationGroupsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListReservationGroups(ctx, req)
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
				},
			},
			{
				Name:  "reservations",
				Usage: "Manage reservations resources",
				Commands: []*cli.Command{

					{
						Name:  "create",
						Usage: "create reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation-id", Usage: "The reservation ID.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.CreateReservationRequest{
								Parent:        parent,
								ReservationId: cmd.String("reservation-id"),
							}

							resp, err := client.CreateReservation(ctx, req)
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
						Usage: "list reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.IntFlag{Name: "page-size", Usage: "The maximum number of items to return per page.", Required: false},
							&cli.StringFlag{Name: "page-token", Usage: "The next_page_token value returned from a previous List request, if any.", Required: false},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.ListReservationsRequest{
								Parent:    parent,
								PageSize:  int32(cmd.Int("page-size")),
								PageToken: cmd.String("page-token"),
							}

							limit := cmd.Int("limit")
							it := client.ListReservations(ctx, req)
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
						Usage: "describe reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.GetReservationRequest{
								Name: name,
							}

							resp, err := client.GetReservation(ctx, req)
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
						Name:  "delete",
						Usage: "delete reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							if !cmd.Bool("quiet") {
								if err := runtime.Confirm(fmt.Sprintf("Execute DeleteReservation on %s?", name)); err != nil {
									return err
								}
							}
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.DeleteReservationRequest{
								Name: name,
							}

							if err := client.DeleteReservation(ctx, req); err != nil {
								return err
							}
							fmt.Println("Done.")
							return nil
						},
					},

					{
						Name:  "update",
						Usage: "update reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							// TODO: nested path field "reservation.name" not yet supported.
							reservation_name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							fmt.Printf("Executing update on %s\n", reservation_name)
							return nil
						},
					},

					{
						Name:  "failover-reservation",
						Usage: "failover-reservation reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "failover-mode", Usage: "A parameter that determines how writes that are pending.", Required: false},
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.FailoverReservationRequest{
								Name:         name,
								FailoverMode: reservationpb.FailoverMode(reservationpb.FailoverMode_value[cmd.String("failover-mode")]),
							}

							resp, err := client.FailoverReservation(ctx, req)
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
						Name:  "describe",
						Usage: "describe reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.GetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.SetIamPolicyRequest{
								Resource: resource,
							}

							resp, err := client.SetIamPolicy(ctx, req)
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
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The Cloud location for the resource.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
							&cli.StringFlag{Name: "project", Usage: "The Cloud project for the resource.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The ID of the reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							resource := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							opts, err := auth.ClientOptions(ctx, cmd)
							if err != nil {
								return err
							}
							client, err := reservation.NewClient(ctx, opts...)
							if err != nil {
								return err
							}
							defer client.Close()

							req := &reservationpb.TestIamPermissionsRequest{
								Resource:    resource,
								Permissions: cmd.StringSlice("permissions"),
							}

							resp, err := client.TestIamPermissions(ctx, req)
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
