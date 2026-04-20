package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/fraenky8/tables-to-go/v2/internal/cmd"
)

var (
	revision       = "master"
	versionTag     = ""
	buildTimestamp = ""
)

func main() {
	ctx := context.Background()

	c := cmd.New(cmd.VersionInfo{
		Revision:       revision,
		VersionTag:     versionTag,
		BuildTimestamp: buildTimestamp,
	}, nil)

	err := c.Run(ctx, os.Args, os.Stdout, os.Stderr)
	if err != nil {
		if !errors.Is(err, cmd.ErrFlagParse) {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
