package simulator

import "context"

var CmdPppoeAccessConcentrator = &CommandSpec{
	[]Token{TPppoe, TAccess, TConcentrator}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoPppoeAccessConcentrator = &CommandSpec{
	[]Token{TNo, TPppoe, TAccess, TConcentrator}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
