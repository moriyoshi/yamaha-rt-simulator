package simulator

import "context"

var CmdPppoePassThroughMember = &CommandSpec{
	[]Token{TPppoe, TPassThrough, TMember}, 35,
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

var CmdNoPppoePassThroughMember = &CommandSpec{
	[]Token{TNo, TPppoe, TPassThrough, TMember}, 4,
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
