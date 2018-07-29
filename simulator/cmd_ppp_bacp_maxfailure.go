package simulator

import "context"

var CmdPppBacpMaxfailure = &CommandSpec{
	[]Token{TPpp, TBacp, TMaxfailure}, 3,
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

var CmdNoPppBacpMaxfailure = &CommandSpec{
	[]Token{TNo, TPpp, TBacp, TMaxfailure}, 4,
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
