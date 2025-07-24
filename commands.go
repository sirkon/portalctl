package main

import (
	"go/ast"
	"go/parser"

	"github.com/sirkon/errors"
)

// Commands структура команд приложения.
type Commands struct {
	Here   CommandHere   `cmd:"" help:"Add current directory to portals under the given name."`
	Delete CommandDelete `cmd:"" help:"Remove portal with the given name."`
	Show   CommandShow   `cmd:"" help:"Show path of the given portal."`
	List   CommandList   `cmd:"" help:"List portals."`
	Prefix CommandPrefix `cmd:"" help:"List all portals with the given prefix."`

	LogCompact CommandLogCompact `cmd:"" help:"Compact existing op log." name:"log-compact"`
	Version    CommandVersion    `cmd:"" help:"Show version and exit."`
	Setup      CommandSetup      `cmd:"" help:"Setup prerequisites."`
}

// RunContext контект исполнения команд.
type RunContext struct {
	cmd *Commands

	appCacheRoot string
	opLogFile    string
}

type identifierName string

// UnmarshalText для реализации encoding.TextUnmarshaler.
func (id *identifierName) UnmarshalText(text []byte) error {
	t := string(text)
	expr, err := parser.ParseExpr(t)
	if err != nil {
		return errors.New("invalid identifier")
	}

	_, ok := expr.(*ast.Ident)
	if !ok {
		return errors.New("invalid identifier")
	}

	(*id) = identifierName(t)
	return nil
}
