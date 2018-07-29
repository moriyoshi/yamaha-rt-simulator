package simulator

import (
	"context"
	"io"
)

var CmdQuit = &CommandSpec{
	[]TokensArityPair{
		{[]Token{TQuit}, 1},
		{[]Token{TExit}, 1},
	}, 0,
	func(_ context.Context, sess *SimulatorSession, _ []TokenInstance) error {
		if sess.Enabled {
			sess.Enabled = false
			return nil
		} else {
			return io.EOF
		}
	},
}
