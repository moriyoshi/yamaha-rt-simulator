package simulator

import "context"

var CmdQacTmWarningUrl = &CommandSpec{
	[]Token{TQacTm, TWarning, TUrl}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQacTmWarningUrl = &CommandSpec{
	[]Token{TNo, TQacTm, TWarning, TUrl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
