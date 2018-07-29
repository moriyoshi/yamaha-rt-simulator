package simulator

import "context"

var CmdQacTmPort = &CommandSpec{
	[]Token{TQacTm, TPort}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmPort = &CommandSpec{
	[]Token{TNo, TQacTm, TPort}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

