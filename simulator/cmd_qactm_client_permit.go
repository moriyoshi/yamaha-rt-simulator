package simulator

import "context"

var CmdQacTmClientPermit = &CommandSpec{
	[]Token{TQacTm, TClient, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmClientPermit = &CommandSpec{
	[]Token{TNo, TQacTm, TClient, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

