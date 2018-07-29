package simulator

import "context"

var CmdQacTmClientPort = &CommandSpec{
	[]Token{TQacTm, TClientPort}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmClientPort = &CommandSpec{
	[]Token{TNo, TQacTm, TClientPort}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
