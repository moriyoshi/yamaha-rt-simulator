package simulator

import (
	"context"
	"strconv"
)

var CmdPPSelect = &CommandSpec{
	[]Token{TPp, TSelect}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if len(tis) != 3 {
			return WrongNumberOfArgs
		}
		n, err := strconv.Atoi(tis[2].Image)
		if err != nil || n < 0 {
			return IntegerArgRequired
		}
		if n < 1 {
			return ArgOutOfRange
		}
		sess.PP = n
		sess.Tunnel = 0
		sess.Log.Debug("pp select ", n)
		return nil
	},
}

var CmdNoPPSelect = &CommandSpec{
	[]Token{TNo, TPp, TSelect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if len(tis) != 3 && len(tis) != 4 {
			return WrongNumberOfArgs
		}
		sess.PP = 0
		return nil
	},
}
