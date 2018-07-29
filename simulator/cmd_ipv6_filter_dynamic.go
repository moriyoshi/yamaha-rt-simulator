package simulator

import "context"

var CmdIpv6FilterDynamic = &CommandSpec{
	[]Token{TIpv6, TFilter, TDynamic}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6FilterDynamic = &CommandSpec{
	[]Token{TNo, TIpv6, TFilter, TDynamic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
