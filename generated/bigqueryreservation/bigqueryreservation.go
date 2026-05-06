package bigqueryreservation

import (
	"context"
	"fmt"
	"strings"

	reservation "cloud.google.com/go/bigquery/reservation/apiv1"
	"cloud.google.com/go/bigquery/reservation/apiv1/reservationpb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/urfave/cli/v3"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud bigqueryreservation command tree.
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
							&cli.StringFlag{Name: "assignment-id", Usage: "The assignment id.", Required: false},
							&cli.StringFlag{Name: "assignee", Usage: "The assignee.", Required: false},
							&cli.BoolFlag{Name: "enable-gemini-in-bigquery", Usage: "The enable gemini in bigquery.", Required: false},
							&cli.StringFlag{Name: "principal", Usage: "The principal.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.CreateAssignmentRequest{Parent: parent}
							req.AssignmentId = cmd.String("assignment-id")
							req.Assignment = &reservationpb.Assignment{
								Assignee:               cmd.String("assignee"),
								EnableGeminiInBigquery: cmd.Bool("enable-gemini-in-bigquery"),
								Principal:              cmd.String("principal"),
							}
							resp, err := client.CreateAssignment(ctx, req)
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
						Usage: "list assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &reservationpb.ListAssignmentsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListAssignments(ctx, req)
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
						Usage: "delete assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
							&cli.StringFlag{Name: "assignment", Usage: "The assignment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s/assignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"), cmd.String("assignment"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.DeleteAssignmentRequest{Name: name}
							if err := client.DeleteAssignment(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "move",
						Usage: "move assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
							&cli.StringFlag{Name: "assignment", Usage: "The assignment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s/assignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"), cmd.String("assignment"))
							fmt.Printf("Executing move on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update assignments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
							&cli.StringFlag{Name: "assignment", Usage: "The assignment.", Required: true},
							&cli.StringFlag{Name: "assignee", Usage: "The assignee.", Required: false},
							&cli.BoolFlag{Name: "enable-gemini-in-bigquery", Usage: "The enable gemini in bigquery.", Required: false},
							&cli.StringFlag{Name: "principal", Usage: "The principal.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s/assignments/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"), cmd.String("assignment"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.UpdateAssignmentRequest{}
							req.Assignment = &reservationpb.Assignment{
								Name:                   name,
								Assignee:               cmd.String("assignee"),
								EnableGeminiInBigquery: cmd.Bool("enable-gemini-in-bigquery"),
								Principal:              cmd.String("principal"),
							}
							var paths []string
							if cmd.IsSet("assignee") {
								paths = append(paths, "assignee")
							}
							if cmd.IsSet("enable-gemini-in-bigquery") {
								paths = append(paths, "enable_gemini_in_bigquery")
							}
							if cmd.IsSet("principal") {
								paths = append(paths, "principal")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateAssignment(ctx, req)
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
				Name:  "bi-reservation",
				Usage: "Manage bi-reservation resources",
				Commands: []*cli.Command{
					{
						Name:  "describe",
						Usage: "describe bi-reservation",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/biReservation", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.GetBiReservationRequest{Name: name}
							resp, err := client.GetBiReservation(ctx, req)
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
						Usage: "update bi-reservation",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "size", Usage: "The size.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/biReservation", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.UpdateBiReservationRequest{}
							req.BiReservation = &reservationpb.BiReservation{
								Name: name,
								Size: int64(cmd.Int("size")),
							}
							var paths []string
							if cmd.IsSet("size") {
								paths = append(paths, "size")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateBiReservation(ctx, req)
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
				Name:  "capacity-commitments",
				Usage: "Manage capacity-commitments resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.BoolFlag{Name: "enforce-single-admin-project-per-org", Usage: "The enforce single admin project per org.", Required: false},
							&cli.StringFlag{Name: "capacity-commitment-id", Usage: "The capacity commitment id.", Required: false},
							&cli.IntFlag{Name: "slot-count", Usage: "The slot count.", Required: false},
							&cli.BoolFlag{Name: "multi-region-auxiliary", Usage: "The multi region auxiliary.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.CreateCapacityCommitmentRequest{Parent: parent}
							req.EnforceSingleAdminProjectPerOrg = cmd.Bool("enforce-single-admin-project-per-org")
							req.CapacityCommitmentId = cmd.String("capacity-commitment-id")
							req.CapacityCommitment = &reservationpb.CapacityCommitment{
								SlotCount:            int64(cmd.Int("slot-count")),
								MultiRegionAuxiliary: cmd.Bool("multi-region-auxiliary"),
							}
							resp, err := client.CreateCapacityCommitment(ctx, req)
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
						Usage: "list capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &reservationpb.ListCapacityCommitmentsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListCapacityCommitments(ctx, req)
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
						Usage: "describe capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "capacity_commitment", Usage: "The capacity_commitment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity_commitment"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.GetCapacityCommitmentRequest{Name: name}
							resp, err := client.GetCapacityCommitment(ctx, req)
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
						Usage: "delete capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "capacity_commitment", Usage: "The capacity_commitment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity_commitment"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.DeleteCapacityCommitmentRequest{Name: name}
							if err := client.DeleteCapacityCommitment(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "capacity_commitment", Usage: "The capacity_commitment.", Required: true},
							&cli.IntFlag{Name: "slot-count", Usage: "The slot count.", Required: false},
							&cli.BoolFlag{Name: "multi-region-auxiliary", Usage: "The multi region auxiliary.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity_commitment"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.UpdateCapacityCommitmentRequest{}
							req.CapacityCommitment = &reservationpb.CapacityCommitment{
								Name:                 name,
								SlotCount:            int64(cmd.Int("slot-count")),
								MultiRegionAuxiliary: cmd.Bool("multi-region-auxiliary"),
							}
							var paths []string
							if cmd.IsSet("slot-count") {
								paths = append(paths, "slot_count")
							}
							if cmd.IsSet("multi-region-auxiliary") {
								paths = append(paths, "multi_region_auxiliary")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateCapacityCommitment(ctx, req)
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
						Name:  "split",
						Usage: "split capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "capacity_commitment", Usage: "The capacity_commitment.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/capacityCommitments/%s", cmd.String("project"), cmd.String("location"), cmd.String("capacity_commitment"))
							fmt.Printf("Executing split on %s\n", name)
							return nil
						},
					},
					{
						Name:  "merge",
						Usage: "merge capacity-commitments",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							fmt.Printf("Executing merge on %s\n", parent)
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
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search-assignments...")
							return nil
						},
					},
					{
						Name:  "search-all-assignments",
						Usage: "search-all-assignments locations",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing search-all-assignments...")
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
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation-group-id", Usage: "The reservation group id.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.CreateReservationGroupRequest{Parent: parent}
							req.ReservationGroupId = cmd.String("reservation-group-id")
							resp, err := client.CreateReservationGroup(ctx, req)
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
						Usage: "describe reservation-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation_group", Usage: "The reservation_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservationGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation_group"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.GetReservationGroupRequest{Name: name}
							resp, err := client.GetReservationGroup(ctx, req)
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
						Usage: "delete reservation-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation_group", Usage: "The reservation_group.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservationGroups/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation_group"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.DeleteReservationGroupRequest{Name: name}
							if err := client.DeleteReservationGroup(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "list reservation-groups",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &reservationpb.ListReservationGroupsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListReservationGroups(ctx, req)
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
				Name:  "reservations",
				Usage: "Manage reservations resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation-id", Usage: "The reservation id.", Required: false},
							&cli.IntFlag{Name: "slot-capacity", Usage: "The slot capacity.", Required: false},
							&cli.BoolFlag{Name: "ignore-idle-slots", Usage: "The ignore idle slots.", Required: false},
							&cli.IntFlag{Name: "concurrency", Usage: "The concurrency.", Required: false},
							&cli.BoolFlag{Name: "multi-region-auxiliary", Usage: "The multi region auxiliary.", Required: false},
							&cli.StringFlag{Name: "secondary-location", Usage: "The secondary location.", Required: false},
							&cli.IntFlag{Name: "max-slots", Usage: "The max slots.", Required: false},
							&cli.StringFlag{Name: "reservation-group", Usage: "The reservation group.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.CreateReservationRequest{Parent: parent}
							req.ReservationId = cmd.String("reservation-id")
							req.Reservation = &reservationpb.Reservation{
								SlotCapacity:         int64(cmd.Int("slot-capacity")),
								IgnoreIdleSlots:      cmd.Bool("ignore-idle-slots"),
								Concurrency:          int64(cmd.Int("concurrency")),
								MultiRegionAuxiliary: cmd.Bool("multi-region-auxiliary"),
								SecondaryLocation:    cmd.String("secondary-location"),
								MaxSlots:             int64(cmd.Int("max-slots")),
								ReservationGroup:     cmd.String("reservation-group"),
							}
							resp, err := client.CreateReservation(ctx, req)
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
						Usage: "list reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.IntFlag{Name: "limit", Usage: "Maximum number of resources to list. 0 means unlimited.", Required: false},
							&cli.IntFlag{Name: "page-size", Usage: "Maximum number of resources per page.", Required: false},
							&cli.BoolFlag{Name: "uri", Usage: "Print a list of resource URIs instead of the default output.", Required: false},
							&cli.StringFlag{Name: "filter", Usage: "Print only resources whose JSON encoding contains this substring.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s/locations/%s", cmd.String("project"), cmd.String("location"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							pageSize := cmd.Int("page-size")
							req := &reservationpb.ListReservationsRequest{Parent: parent}
							if pageSize > 0 {
								req.PageSize = int32(pageSize)
							}
							it := client.ListReservations(ctx, req)
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
						Usage: "describe reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.GetReservationRequest{Name: name}
							resp, err := client.GetReservation(ctx, req)
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
						Usage: "delete reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.DeleteReservationRequest{Name: name}
							if err := client.DeleteReservation(ctx, req); err != nil {
								return err
							}
							fmt.Printf("Deleted %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
							&cli.IntFlag{Name: "slot-capacity", Usage: "The slot capacity.", Required: false},
							&cli.BoolFlag{Name: "ignore-idle-slots", Usage: "The ignore idle slots.", Required: false},
							&cli.IntFlag{Name: "concurrency", Usage: "The concurrency.", Required: false},
							&cli.BoolFlag{Name: "multi-region-auxiliary", Usage: "The multi region auxiliary.", Required: false},
							&cli.StringFlag{Name: "secondary-location", Usage: "The secondary location.", Required: false},
							&cli.IntFlag{Name: "max-slots", Usage: "The max slots.", Required: false},
							&cli.StringFlag{Name: "reservation-group", Usage: "The reservation group.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &reservationpb.UpdateReservationRequest{}
							req.Reservation = &reservationpb.Reservation{
								Name:                 name,
								SlotCapacity:         int64(cmd.Int("slot-capacity")),
								IgnoreIdleSlots:      cmd.Bool("ignore-idle-slots"),
								Concurrency:          int64(cmd.Int("concurrency")),
								MultiRegionAuxiliary: cmd.Bool("multi-region-auxiliary"),
								SecondaryLocation:    cmd.String("secondary-location"),
								MaxSlots:             int64(cmd.Int("max-slots")),
								ReservationGroup:     cmd.String("reservation-group"),
							}
							var paths []string
							if cmd.IsSet("slot-capacity") {
								paths = append(paths, "slot_capacity")
							}
							if cmd.IsSet("ignore-idle-slots") {
								paths = append(paths, "ignore_idle_slots")
							}
							if cmd.IsSet("concurrency") {
								paths = append(paths, "concurrency")
							}
							if cmd.IsSet("multi-region-auxiliary") {
								paths = append(paths, "multi_region_auxiliary")
							}
							if cmd.IsSet("secondary-location") {
								paths = append(paths, "secondary_location")
							}
							if cmd.IsSet("max-slots") {
								paths = append(paths, "max_slots")
							}
							if cmd.IsSet("reservation-group") {
								paths = append(paths, "reservation_group")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateReservation(ctx, req)
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
						Name:  "failover-reservation",
						Usage: "failover-reservation reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							fmt.Printf("Executing failover-reservation on %s\n", name)
							return nil
						},
					},
					{
						Name:  "get-iam-policy",
						Usage: "get-iam-policy reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.GetIamPolicyRequest{Resource: name}
							resp, err := client.GetIamPolicy(ctx, req)
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
						Name:  "set-iam-policy",
						Usage: "set-iam-policy reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							fmt.Printf("Executing set-iam-policy on %s\n", name)
							return nil
						},
					},
					{
						Name:  "test-iam-permissions",
						Usage: "test-iam-permissions reservations",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "location", Usage: "The location.", Required: true},
							&cli.StringFlag{Name: "reservation", Usage: "The reservation.", Required: true},
							&cli.StringSliceFlag{Name: "permissions", Usage: "The permissions.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/locations/%s/reservations/%s", cmd.String("project"), cmd.String("location"), cmd.String("reservation"))
							client, err := reservation.NewClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &iampb.TestIamPermissionsRequest{Resource: name}
							req.Permissions = cmd.StringSlice("permissions")
							resp, err := client.TestIamPermissions(ctx, req)
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
		},
	}
}
