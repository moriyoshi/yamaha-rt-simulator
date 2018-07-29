package simulator

import "context"

var CmdPppMscbcpMaxretry = &CommandSpec{
	[]Token{TPpp, TMscbcp, TMaxretry}, 3,
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

var CmdNoPppMscbcpMaxretry = &CommandSpec{
	[]Token{TNo, TPpp, TMscbcp, TMaxretry}, 4,
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
