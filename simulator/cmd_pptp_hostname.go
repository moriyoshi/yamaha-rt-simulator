package simulator

import "context"

var CmdPptpHostname = &CommandSpec{
	[]Token{TPptp, THostname}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpHostname = &CommandSpec{
	[]Token{TNo, TPptp, THostname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
