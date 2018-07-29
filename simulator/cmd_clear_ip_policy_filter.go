package simulator

import "context"

var CmdClearIpPolicyFilter = &CommandSpec{
	[]Token{TClear, TIp, TPolicy, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdClearIpv6PolicyFilter = &CommandSpec{
	[]Token{TClear, TIpv6, TPolicy, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

