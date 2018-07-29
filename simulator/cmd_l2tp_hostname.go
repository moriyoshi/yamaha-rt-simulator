package simulator

import "context"

var CmdL2TpHostname = &CommandSpec{
	[]Token{TL2Tp, THostname}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpHostname = &CommandSpec{
	[]Token{TNo, TL2Tp, THostname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

