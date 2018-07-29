package simulator

import "context"

var CmdSipOuterAddress = &CommandSpec{
	[]Token{TSip, TOuter, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipOuterAddress = &CommandSpec{
	[]Token{TNo, TSip, TOuter, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

