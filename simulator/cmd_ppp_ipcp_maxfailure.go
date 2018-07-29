package simulator

import "context"

var CmdPppIpcpMaxfailure = &CommandSpec{
	[]Token{TPpp, TIpcp, TMaxfailure}, 3,
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

var CmdNoPppIpcpMaxfailure = &CommandSpec{
	[]Token{TNo, TPpp, TIpcp, TMaxfailure}, 4,
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
