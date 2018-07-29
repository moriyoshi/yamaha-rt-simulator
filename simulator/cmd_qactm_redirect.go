package simulator

import "context"

var CmdQacTmRedirect = &CommandSpec{
	[]Token{TQacTm, TRedirect}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmRedirect = &CommandSpec{
	[]Token{TNo, TQacTm, TRedirect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

