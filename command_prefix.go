package main

import (
	"fmt"
	"os"
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
			fmt.Println(portal + "/")
		}

		return nil
	}

	if pos == 0 {
		return nil
	}

	parts := strings.Split(d.Prefix, "/")
	if len(parts) != 2 {
		return nil
	}

	path, err := portallog.LogShowPortalPath(ctx.opLogFile, parts[0])
	if err != nil {
		return nil
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		if !strings.HasPrefix(dir.Name(), parts[1]) {
			continue
		}

		fmt.Println(filepath.Join(parts[0], dir.Name()))
	}

	return nil
}
