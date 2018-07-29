package simulator

import "context"

var CmdPptpWindowSize = &CommandSpec{
	[]Token{TPptp, TWindow, TSize}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpWindowSize = &CommandSpec{
	[]Token{TNo, TPptp, TWindow, TSize}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
