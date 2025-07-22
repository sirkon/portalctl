package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirkon/errors"
	"github.com/sirkon/portalctl/internal/portallog"
)

// CommandShow реализация команды show.
type CommandShow struct {
	Name string `arg:"" required:"true" help:"Portal name."`
}

// Run запуск команды.
func (d CommandShow) Run(ctx *RunContext) error {
	pos := strings.IndexByte(d.Name, '/')
	if pos == -1 {
		path, err := portallog.LogShowPortalPath(ctx.opLogFile, string(d.Name))
		if err != nil {
			return errors.Wrapf(err, "look for portal %q", d.Name)
		}

		fmt.Println(path)
		return nil
	}
	if pos == 0 {
		return nil
	}

	path, err := portallog.LogShowPortalPath(ctx.opLogFile, d.Name[:pos])
	if err != nil {
		return errors.Wrapf(err, "look for portal of %q path", d.Name)
	}

	suffix := d.Name[pos+1:]

	fmt.Println(filepath.Join(path, suffix))
	return nil
}
