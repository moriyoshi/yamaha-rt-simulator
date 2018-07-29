package simulator

import "context"

var CmdIpPolicyFilter = &CommandSpec{
	[]Token{TIp, TPolicy, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PolicyFilter = &CommandSpec{
	[]Token{TIpv6, TPolicy, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPolicyFilter = &CommandSpec{
	[]Token{TNo, TIp, TPolicy, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PolicyFilter = &CommandSpec{
	[]Token{TNo, TIpv6, TPolicy, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

