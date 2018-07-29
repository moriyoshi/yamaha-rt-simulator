package simulator

import "context"

var CmdPppIpcpMaxconfigure = &CommandSpec{
	[]Token{TPpp, TIpcp, TMaxconfigure}, 3,
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

var CmdNoPppIpcpMaxconfigure = &CommandSpec{
	[]Token{TNo, TPpp, TIpcp, TMaxconfigure}, 4,
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
