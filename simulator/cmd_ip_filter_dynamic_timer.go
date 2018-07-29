package simulator

import "context"

var CmdIpFilterDynamicTimer = &CommandSpec{
	[]Token{TIp, TFilter, TDynamic, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFilterDynamicTimer = &CommandSpec{
	[]Token{TNo, TIp, TFilter, TDynamic, TTimer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

