package connectors

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

// Command returns the gcloud connectors command tree.
func Command() *cli.Command {
	return &cli.Command{
		Name:     "connectors",
		Usage:    "manage AlloyDB connectors resources",
		Commands: []*cli.Command{},
	}
}
