package simulator

import "context"

var CmdIpForwardFilter = &CommandSpec{
	[]Token{TIp, TForward, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpForwardFilter = &CommandSpec{
	[]Token{TNo, TIp, TForward, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

