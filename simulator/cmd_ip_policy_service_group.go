package simulator

import "context"

var CmdIpPolicyServiceGroup = &CommandSpec{
	[]Token{TIp, TPolicy, TService, TGroup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PolicyServiceGroup = &CommandSpec{
	[]Token{TIpv6, TPolicy, TService, TGroup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPolicyServiceGroup = &CommandSpec{
	[]Token{TNo, TIp, TPolicy, TService, TGroup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PolicyServiceGroup = &CommandSpec{
	[]Token{TNo, TIpv6, TPolicy, TService, TGroup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

