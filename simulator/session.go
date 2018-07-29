package simulator

import (
	"context"
	"fmt"
	"io"
	"sort"
	"text/template"
	"time"

	base "github.com/moriyoshi/yamaha-rt-simulator"
	"golang.org/x/text/encoding"
)

type SimulatorSession struct {
	*Simulator
	Dirty    bool
	Encoder  *encoding.Encoder
	Terminal base.Terminal
	Charset  NLSInfo
	Prompt   string
	Enabled  bool
	PP       int
	Tunnel   int
}

const (
	welcomeMessage = `{{.ModelName}} Rev.{{.Version}} ({{formatAsRFC822 .Date}})
{{.Copyright}}
To display the software copyright statement, use 'show copyright' command.
{{range $index, $:= .PhyInterfaces}}{{if gt $index 0}}, {{end}}{{.HardwareAddr}}{{end}}
Memory 256Mbytes, {{len .PhyInterfaces}}LAN, 1BRI
`
)

var greetTemplate *template.Template

func formatAsRFC822(t time.Time) string {
	return t.Format("Mon Jan _2 15:04:05 2006")
}

func parseIntoTokenImages(cmd string) (tokens []string, endsWithSpaces bool) {
	im := []byte{}
	endsWithSpaces = true
	for i := 0; i < len(cmd); {
		c := cmd[i]
		switch c {
		case ' ', '\t':
			tokens = append(tokens, string(im))
			im = []byte{}
			i++
			for ; i < len(cmd); i++ {
				switch cmd[i] {
				case ' ', '\t':
					continue
				}
				break
			}
		case '"', '\'':
			d := c
			i++
			for {
				if i >= len(cmd) {
					break
				}
				c := cmd[i]
				if c == d {
					i++
					break
				} else if c == '\\' {
					i++
					if i < len(cmd) {
						im = append(im, cmd[i])
						i++
					}

				} else {
					im = append(im, c)
					i++
				}
			}
		case '\\':
			i++
			if i < len(cmd) {
				im = append(im, cmd[i])
				i++
			}
		default:
			im = append(im, c)
			i++
		}
	}
	if len(im) > 0 {
		tokens = append(tokens, string(im))
		endsWithSpaces = false
	}
	return
}

func (sess *SimulatorSession) ResetEncoder() {
	sess.Encoder = sess.Charset.NewEncoder()
}

func (sess *SimulatorSession) ResetPrompt() {
	sess.Prompt = sess.RunningConfig.Prompt
}

func (sess *SimulatorSession) Reset() {
	sess.ResetEncoder()
	sess.ResetPrompt()
	sess.Enabled = false
	sess.PP = 0
}

func (sess *SimulatorSession) autoCompleteCallback(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key != 0x09 && key != 0x20 {
		return "", 0, false
	}
	ims, endsWithSpaces := parseIntoTokenImages(line)
	tokens, candidates, candidatesForDisplay := sess.CommandTrie.LookupByStringC(sess, ims)
	if len(ims) == len(tokens)+1 && !endsWithSpaces {
		if key != 0x20 && len(candidatesForDisplay) > 0 {
			buf := []byte(sess.GetPrompt())
			buf = append(buf, []byte(line)...)
			buf = append(buf, '?', 0x0a)
			buf = append(buf, '?')
			sort.Strings(candidatesForDisplay)
			for _, c := range candidatesForDisplay {
				buf = append(buf, ' ')
				buf = append(buf, []byte(c)...)
			}
			buf = append(buf, '\n')
			sess.Terminal.Write(buf)
			return line, pos, true
		}
	} else if len(ims) == len(tokens) && !endsWithSpaces {
		if len(candidates) == 0 {
			_newLine := []byte{}
			for _, t := range tokens {
				_newLine = append(_newLine, []byte(t.Image)...)
				_newLine = append(_newLine, ' ')
			}
			return string(_newLine), len(_newLine), true
		} else {
			return line + " ", len(line) + 1, true
		}
	}
	return "", 0, false
}

func (sess *SimulatorSession) greet() error {
	return greetTemplate.Execute(sess.Terminal, sess)
}

func (sess *SimulatorSession) getPromptSuffix() string {
	var suffix string
	if sess.PP > 0 {
		suffix += fmt.Sprintf("pp%d", sess.PP)
	}
	if sess.Tunnel > 0 {
		suffix += fmt.Sprintf("tunnel%d", sess.Tunnel)
	}
	if sess.Enabled {
		suffix += "# "
	} else {
		suffix += "> "
	}
	return suffix
}

func (sess *SimulatorSession) GetPrompt() string {
	return sess.Prompt + sess.getPromptSuffix()
}

func (sess *SimulatorSession) parseAndRunCommand(ctx context.Context, line string) error {
	ims, _ := parseIntoTokenImages(line)
	if len(ims) == 0 {
		return nil
	}
	tis, cmdHdlr := sess.CommandTrie.LookupByString(sess, ims)
	if cmdHdlr == nil {
		return InvalidCommandName
	}
	for len(tis) < len(ims) {
		tis = append(tis, TokenInstance{TSomething, ims[len(tis)]})
	}
	canonLine := []byte{}
	for i, ti := range tis {
		if i > 0 {
			canonLine = append(canonLine, ' ')
		}
		canonLine = append(canonLine, []byte(ti.Image)...)
	}
	sess.Log.Debug(fmt.Sprintf("command %s entered", canonLine))
	return cmdHdlr.(*CommandSpec).Handler(ctx, sess, tis)
}

func (sess *SimulatorSession) awaitCommand(ctx context.Context) error {
	sess.Terminal.SetPrompt(sess.GetPrompt())
	line, err := sess.Terminal.ReadLine()
	if err != nil {
		return err
	}
	return sess.parseAndRunCommand(ctx, line)
}

func (sess *SimulatorSession) Write(msg string) (int, error) {
	b, err := sess.Encoder.Bytes([]byte(msg))
	if err != nil {
		return 0, err
	}
	return sess.Terminal.Write(b)
}

func (sess *SimulatorSession) WriteF(format string, args ...interface{}) (int, error) {
	return sess.Write(fmt.Sprintf(format, args...))
}

func (sess *SimulatorSession) GetText(any interface{}) string {
	catalog := messages[sess.Charset.Locale]
	msg, ok := catalog[any]
	if !ok {
		return fmt.Sprintf("%v", any)
	} else {
		return msg
	}
}

func (sess *SimulatorSession) Run(ctx context.Context) {
	sess.greet()
outer:
	for {
		err := sess.awaitCommand(ctx)
		if err != nil {
			switch err {
			case io.EOF:
				break outer
			default:
				sess.Write(fmt.Sprintf(sess.GetText("Error: %s\n"), sess.GetText(err)))
			}
		}
	}
}
