package simulator

import "context"

var CmdPppIpcpMsext = &CommandSpec{
	[]Token{TPpp, TIpcp, TMsext}, 3,
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

var CmdNoPppIpcpMsext = &CommandSpec{
	[]Token{TNo, TPpp, TIpcp, TMsext}, 4,
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
