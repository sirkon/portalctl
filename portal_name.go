package main

import (
	"regexp"

	"github.com/sirkon/errors"
)

// PortalName имя портала.
type PortalName string

// UnmarshalText для реализации decoder.TextUnmarshaler.
func (s *PortalName) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return errors.New("portal name must not be empty")
	}

	if !identifierMatch.Match(data) {
		return errors.Newf("portal name must match [a-zA-Z_][a-zA-Z0-9_]*")
	}

	*s = PortalName(data)
	return nil
}

var (
	identifierMatch = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")
)
