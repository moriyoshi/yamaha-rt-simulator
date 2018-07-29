package simulator

import (
	"context"
	"strconv"
)

var CmdTunnelSelect = &CommandSpec{
	[]Token{TTunnel, TSelect}, 2,
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
		sess.PP = 0
		sess.Tunnel = n
		sess.Log.Debug("tunnel select ", n)
		return nil
	},
}

var CmdNoTunnelSelect = &CommandSpec{
	[]Token{TNo, TTunnel, TSelect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if len(tis) != 3 && len(tis) != 4 {
			return WrongNumberOfArgs
		}
		sess.Tunnel = 0
		return nil
	},
}
