package simulator

import "context"

var CmdPppLcpMaxfailure = &CommandSpec{
	[]Token{TPpp, TLcp, TMaxfailure}, 3,
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

var CmdNoPppLcpMaxfailure = &CommandSpec{
	[]Token{TNo, TPpp, TLcp, TMaxfailure}, 4,
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
