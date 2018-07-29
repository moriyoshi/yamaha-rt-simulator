package simulator

import "context"

var CmdPppBacpMaxconfigure = &CommandSpec{
	[]Token{TPpp, TBacp, TMaxconfigure}, 3,
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

var CmdNoPppBacpMaxconfigure = &CommandSpec{
	[]Token{TNo, TPpp, TBacp, TMaxconfigure}, 4,
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
