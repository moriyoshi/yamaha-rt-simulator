package simulator

import "context"

var CmdPppIpcpMaxterminate = &CommandSpec{
	[]Token{TPpp, TIpcp, TMaxterminate}, 3,
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

var CmdNoPppIpcpMaxterminate = &CommandSpec{
	[]Token{TNo, TPpp, TIpcp, TMaxterminate}, 4,
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
