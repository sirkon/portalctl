package main

import (
	"github.com/sirkon/errors"
	"github.com/sirkon/portalctl/internal/portallog"
)

// CommandDelete реализация команды delete.
type CommandDelete struct {
	Name PortalName `arg:"" required:"" help:"Portal name to delete."`
}

// Run запуск команды.
func (d CommandDelete) Run(ctx *RunContext) error {
	data, err := portallog.LogRead(ctx.opLogFile)
	if err != nil {
		return errors.Wrap(err, "read op log data")
	}

	if _, ok := data[string(d.Name)]; !ok {
		return errors.Wrapf(err, "unknown portal %q", d.Name)
	}

	oplog := portallog.NewLog(ctx.opLogFile)
	if err := oplog.DeletePortal(string(d.Name)); err != nil {
		return errors.Wrap(err, "store data")
	}

	return nil
}
