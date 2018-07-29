package simulator

import "context"

var CmdQacTmVersionMargin = &CommandSpec{
	[]Token{TQacTm, TVersion, TMargin}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmVersionMargin = &CommandSpec{
	[]Token{TNo, TQacTm, TVersion, TMargin}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

