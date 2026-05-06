package osconfig

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud osconfig command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:     "osconfig",
		Usage:    "manage OS Config API resources",
		Commands: []*cli.Command{},
	}
}
