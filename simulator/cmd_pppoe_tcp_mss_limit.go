package simulator

import "context"

var CmdPppoeTcpMssLimit = &CommandSpec{
	[]Token{TPppoe, TTcp, TMss, TLimit}, 4,
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

var CmdNoPppoeTcpMssLimit = &CommandSpec{
	[]Token{TNo, TPppoe, TTcp, TMss, TLimit}, 5,
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
