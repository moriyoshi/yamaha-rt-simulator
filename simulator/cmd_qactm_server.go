package simulator

import "context"

var CmdQacTmServer = &CommandSpec{
	[]Token{TQacTm, TServer}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmServer = &CommandSpec{
	[]Token{TNo, TQacTm, TServer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

