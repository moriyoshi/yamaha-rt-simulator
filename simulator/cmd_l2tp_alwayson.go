package simulator

import "context"

var CmdL2TpAlwaysOn = &CommandSpec{
	[]Token{TL2Tp, TAlwaysOn}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpAlwaysOn = &CommandSpec{
	[]Token{TNo, TL2Tp, TAlwaysOn}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

