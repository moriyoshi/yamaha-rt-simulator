package simulator

import "context"

var CmdPppLcpMaxterminate = &CommandSpec{
	[]Token{TPpp, TLcp, TMaxterminate}, 3,
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

var CmdNoPppLcpMaxterminate = &CommandSpec{
	[]Token{TNo, TPpp, TLcp, TMaxterminate}, 4,
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
