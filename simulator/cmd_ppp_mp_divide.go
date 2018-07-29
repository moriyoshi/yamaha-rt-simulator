package simulator

import "context"

var CmdPppMpDivide = &CommandSpec{
	[]Token{TPpp, TMp, TDivide}, 3,
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

var CmdNoPppMpDivide = &CommandSpec{
	[]Token{TNo, TPpp, TMp, TDivide}, 4,
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
