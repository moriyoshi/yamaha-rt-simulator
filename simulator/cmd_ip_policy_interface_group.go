package simulator

import "context"

var CmdIpPolicyInterfaceGroup = &CommandSpec{
	[]Token{TIp, TPolicy, TInterface, TGroup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PolicyInterfaceGroup = &CommandSpec{
	[]Token{TIpv6, TPolicy, TInterface, TGroup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPolicyInterfaceGroup = &CommandSpec{
	[]Token{TNo, TIp, TPolicy, TInterface, TGroup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PolicyInterfaceGroup = &CommandSpec{
	[]Token{TNo, TIpv6, TPolicy, TInterface, TGroup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

