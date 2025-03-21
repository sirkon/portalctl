package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirkon/errors"
	"github.com/sirkon/message"
)

func bashRCAppend(jumper string) []string {
	return []string{
		jumper + "() {",
		"    cd `portalctl show $1`",
		"}",
		"complete_portal() {",
		"    COMPREPLY=($(compgen -W \"`portalctl prefix`\" \"$2\"))",
		"}",
		"complete -F complete_portal " + jumper,
	}
}

func zshRCAppend(jumper string) []string {
	return []string{
		jumper + "() {",
		"    cd `portalctl show $1`",
		"}",
		"complete -o nospace -C 'portalctl prefix' " + jumper,
	}
}

// CommandSetup реализация команды setup.
type CommandSetup struct {
	JumpFunction identifierName `help:"Name of the jump function." default:"portal" short:"j"`
}

// Run запуск команды.
func (d CommandSetup) Run(ctx *RunContext) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "get user homedir")
	}

	filesToTouch := map[string][]string{
		".bashrc": bashRCAppend(string(d.JumpFunction)),
		".zshrc":  zshRCAppend(string(d.JumpFunction)),
	}
	for file, rcAppend := range filesToTouch {
		if err := installIntoFile(homedir, file, rcAppend); err != nil {
			if !errors.Is(err, justPassThisRC) {
				return errors.Wrap(err, "process "+file)
			}
			continue
		}

		message.Infof("%s done", file)
	}

	return nil
}

const justPassThisRC = errors.Const("this file was not found and it is OK")

func installIntoFile(home, rc string, rcAppend []string) error {
	rcFullPath := filepath.Join(home, rc)
	stat, err := os.Stat(rcFullPath)
	if err != nil {
		if os.IsNotExist(err) {
			message.Infof("%s not found, omitting", rc)
			return justPassThisRC
		}
	}
	if !stat.Mode().IsRegular() {
		return errors.New("is not a regular file")
	}

	data, err := os.ReadFile(rcFullPath)
	if err != nil {
		return errors.Wrap(err, "read file")
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, rcAppend[0]) {
			message.Infof("%s has already been set up, omitting", rc)
			return nil
		}
	}

	// Убираем пустые строки в конце
	for len(lines) > 0 {
		if lines[len(lines)-1] != "" {
			break
		}

		lines = lines[:len(lines)-1]
	}

	var builder bytes.Buffer
	for _, line := range lines {
		builder.WriteString(line)
		builder.WriteByte('\n')
	}
	builder.WriteByte('\n')
	for _, line := range rcAppend {
		builder.WriteString(line)
		builder.WriteByte('\n')
	}
	builder.WriteByte('\n')

	rcTmpPath := filepath.Join(home, ".rc-tmp")
	if err := os.WriteFile(rcTmpPath, builder.Bytes(), 0644); err != nil {
		return errors.Wrap(err, "build a temporary file with prerequisites")
	}

	if err := os.Rename(rcTmpPath, rcFullPath); err != nil {
		return errors.Wrap(err, "replace original with the temporary file")
	}

	return nil
}
