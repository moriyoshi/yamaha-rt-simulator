package simulator

import "context"

var CmdIpFilterDynamic = &CommandSpec{
	[]Token{TIp, TFilter, TDynamic}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFilterDynamic = &CommandSpec{
	[]Token{TNo, TIp, TFilter, TDynamic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
