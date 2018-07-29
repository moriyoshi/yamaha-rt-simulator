package yamaha_rt_simulator

import (
	"context"
)

type Terminal interface {
	ReadLine() (string, error)
	ReadPassword(string) (string, error)
	Write([]byte) (int, error)
	SetPrompt(string)
}

type AutoCompleteCallback func(string, int, rune) (string, int, bool)

type TerminalFactory func(AutoCompleteCallback) Terminal

type ApplianceSession interface {
	Run(context.Context)
}

type Appliance interface {
	NewSession(TerminalFactory) ApplianceSession
}
