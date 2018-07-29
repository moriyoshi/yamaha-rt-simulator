package simulator

import "context"

var CmdIpFilterSourceRoute = &CommandSpec{
	[]Token{TIp, TFilter, TSourceRoute}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFilterSourceRoute = &CommandSpec{
	[]Token{TNo, TIp, TFilter, TSourceRoute}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

