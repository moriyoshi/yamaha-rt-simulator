package simulator

import "context"

var CmdPppCcpMaxterminate = &CommandSpec{
	[]Token{TPpp, TCcp, TMaxterminate}, 3,
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

var CmdNoPppCcpMaxterminate = &CommandSpec{
	[]Token{TNo, TPpp, TCcp, TMaxterminate}, 4,
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
