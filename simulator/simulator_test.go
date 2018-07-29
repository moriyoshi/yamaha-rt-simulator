package simulator

import (
	"bytes"
	"testing"

	base "github.com/moriyoshi/yamaha-rt-simulator"
	"github.com/stretchr/testify/assert"
)

func TestParseIntoTokenImages(t *testing.T) {
	type Case struct {
		Input          string
		Result         []string
		EndsWithSpaces bool
	}
	cases := []Case{
		Case{"show config", []string{"show", "config"}, false},
		Case{"show config ", []string{"show", "config"}, true},
		Case{"show\tconfig", []string{"show", "config"}, false},
		Case{"show  config", []string{"show", "config"}, false},
		Case{"show \t \tconfig", []string{"show", "config"}, false},
		Case{"sh\\ow config", []string{"show", "config"}, false},
		Case{"'sh\\ow' config", []string{"show", "config"}, false},
		Case{"'sh\\ow' config", []string{"show", "config"}, false},
		Case{"\"show\" config", []string{"show", "config"}, false},
		Case{"\"sh\\ow\" config", []string{"show", "config"}, false},
		Case{"'sh'ow config", []string{"show", "config"}, false},
		Case{"\"sh\"ow config", []string{"show", "config"}, false},
		Case{"\"sh\"'ow' config", []string{"show", "config"}, false},
	}

	for _, c := range cases {
		t.Run(c.Input, func(t *testing.T) {
			tokens, endsWithSpaces := parseIntoTokenImages(c.Input)
			assert.Equal(t, c.Result, tokens)
			assert.Equal(t, c.EndsWithSpaces, endsWithSpaces)
		})
	}
}

type dummyTerminal struct {
	Line   string
	Prompt string
	B      bytes.Buffer
}

func (dt *dummyTerminal) ReadLine() (string, error) {
	return dt.Line, nil
}

func (dt *dummyTerminal) ReadPassword(string) (string, error) {
	return dt.Line, nil
}

func (dt *dummyTerminal) SetPrompt(v string) {
	dt.Prompt = v
}

func (dt *dummyTerminal) Write(b []byte) (int, error) {
	return dt.B.Write(b)
}

type dummyLogger struct{}

func (*dummyLogger) Debug(...interface{}) {}
func (*dummyLogger) Info(...interface{})  {}
func (*dummyLogger) Error(...interface{}) {}

func TestAutoCompleteCallback(t *testing.T) {
	type Case struct {
		Input   string
		Pos     int
		Rune    rune
		Buf     []byte
		NewLine string
		NewPos  int
		Ok      bool
	}
	cases := []Case{
		Case{"show", 0, 9, nil, "show ", 5, true},
		Case{"show co", 0, 9, []byte("> show co?\n? config copyright\n"), "show co", 0, true},
	}

	for _, c := range cases {
		t.Run(c.Input, func(t *testing.T) {
			dt := &dummyTerminal{}
			sim := NewSimulator(&Spec{}, &SimulatorConfig{Charset: NLSInfo{"default", "US-ASCII", "en"}}, &dummyLogger{})
			sess := sim.NewSession(func(base.AutoCompleteCallback) base.Terminal { return dt })
			newLine, newPos, ok := sess.(*SimulatorSession).autoCompleteCallback(c.Input, c.Pos, c.Rune)
			assert.Equal(t, c.Buf, dt.B.Bytes())
			assert.Equal(t, c.NewLine, newLine)
			assert.Equal(t, c.NewPos, newPos)
			assert.Equal(t, c.Ok, ok)
		})
	}
}
