package simulator

import "context"

var CmdIpPolicyAddressGroup = &CommandSpec{
	[]Token{TIp, TPolicy, TAddress, TGroup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PolicyAddressGroup = &CommandSpec{
	[]Token{TIpv6, TPolicy, TAddress, TGroup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPolicyAddressGroup = &CommandSpec{
	[]Token{TNo, TIp, TPolicy, TAddress, TGroup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PolicyAddressGroup = &CommandSpec{
	[]Token{TNo, TIpv6, TPolicy, TAddress, TGroup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

