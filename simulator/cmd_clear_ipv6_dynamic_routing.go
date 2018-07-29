package simulator

import "context"

var CmdClearIpv6DynamicRouting = &CommandSpec{
	[]Token{TClear, TIpv6, TDynamic, TRouting}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

