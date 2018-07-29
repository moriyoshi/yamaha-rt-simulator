package simulator

import "context"

var CmdQacTmClientUpdate = &CommandSpec{
	[]Token{TQacTm, TClient, TUpdate}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmClientUpdate = &CommandSpec{
	[]Token{TNo, TQacTm, TClient, TUpdate}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

