package devicestreaming

import (
	"context"
	"fmt"

	devicestreaming "cloud.google.com/go/devicestreaming/apiv1"
	"cloud.google.com/go/devicestreaming/apiv1/devicestreamingpb"
	"github.com/urfave/cli/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Command returns the gcloud devicestreaming command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "devicestreaming",
		Usage: "manage Device Streaming API resources",
		Commands: []*cli.Command{
			{
				Name:  "device-sessions",
				Usage: "Manage device-sessions resources",
				Commands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device-session-id", Usage: "The device session id.", Required: false},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							parent := fmt.Sprintf("projects/%s", cmd.String("project"))
							client, err := devicestreaming.NewDirectAccessClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &devicestreamingpb.CreateDeviceSessionRequest{Parent: parent}
							req.DeviceSessionId = cmd.String("device-session-id")
							req.DeviceSession = &devicestreamingpb.DeviceSession{
								Name: cmd.String("name"),
							}
							resp, err := client.CreateDeviceSession(ctx, req)
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
						Usage: "list device-sessions",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							fmt.Println("Executing list...")
							return nil
						},
					},
					{
						Name:  "describe",
						Usage: "describe device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device_session", Usage: "The device_session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deviceSessions/%s", cmd.String("project"), cmd.String("device_session"))
							client, err := devicestreaming.NewDirectAccessClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &devicestreamingpb.GetDeviceSessionRequest{Name: name}
							resp, err := client.GetDeviceSession(ctx, req)
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
						Usage: "cancel device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device_session", Usage: "The device_session.", Required: true},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deviceSessions/%s", cmd.String("project"), cmd.String("device_session"))
							fmt.Printf("Executing cancel on %s\n", name)
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update device-sessions",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "device_session", Usage: "The device_session.", Required: true},
							&cli.StringFlag{Name: "name", Usage: "The name.", Required: false},
						},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							name := fmt.Sprintf("projects/%s/deviceSessions/%s", cmd.String("project"), cmd.String("device_session"))
							client, err := devicestreaming.NewDirectAccessClient(ctx)
							if err != nil {
								return err
							}
							defer client.Close()
							req := &devicestreamingpb.UpdateDeviceSessionRequest{}
							req.DeviceSession = &devicestreamingpb.DeviceSession{
								Name: name,
								Name: cmd.String("name"),
							}
							var paths []string
							if cmd.IsSet("name") {
								paths = append(paths, "name")
							}
							req.UpdateMask = &fieldmaskpb.FieldMask{Paths: paths}
							resp, err := client.UpdateDeviceSession(ctx, req)
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
