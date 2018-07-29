package simulator

import "context"

var CmdClearIpDynamicRouting = &CommandSpec{
	[]Token{TClear, TIp, TDynamic, TRouting}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

