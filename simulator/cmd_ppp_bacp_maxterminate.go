package simulator

import "context"

var CmdPppBacpMaxterminate = &CommandSpec{
	[]Token{TPpp, TBacp, TMaxterminate}, 3,
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

var CmdNoPppBacpMaxterminate = &CommandSpec{
	[]Token{TNo, TPpp, TBacp, TMaxterminate}, 4,
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
