package simulator

import "context"

var CmdBgpForceToAdvertise = &CommandSpec{
	[]Token{TBgp, TForceToAdvertise}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpForceToAdvertise = &CommandSpec{
	[]Token{TNo, TBgp, TForceToAdvertise}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

