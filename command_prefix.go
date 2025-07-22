package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirkon/portalctl/internal/portallog"
)

// CommandPrefix реализация команды prefix.
type CommandPrefix struct {
	Prefix string `arg:"" help:"Portal name prefix." default:""`
}

// Run запуск команды.
func (d CommandPrefix) Run(ctx *RunContext) error {
	pos := strings.Index(d.Prefix, "/")
	if pos < 0 {
		portals, err := portallog.LogFilter(ctx.opLogFile, d.Prefix)
		if err != nil {
			return nil
		}

		for _, portal := range portals {
			fmt.Println(portal)
		}

		return nil
	}

	if pos == 0 {
		return nil
	}

	path, err := portallog.LogShowPortalPath(ctx.opLogFile, d.Prefix[:pos])
	if err != nil {
		return nil
	}

	matches, err := filepath.Glob(filepath.Join(path, d.Prefix[pos+1:]+"*"))
	if err != nil {
		return nil
	}

	for _, match := range matches {
		fmt.Println(d.Prefix[:pos] + strings.TrimPrefix(match, path))
	}

	return nil
}
