package simulator

import "context"

var CmdIpPolicyFilterSet = &CommandSpec{
	[]Token{TIp, TPolicy, TFilter, TSet}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PolicyFilterSet = &CommandSpec{
	[]Token{TIpv6, TPolicy, TFilter, TSet}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPolicyFilterSet = &CommandSpec{
	[]Token{TNo, TIp, TPolicy, TFilter, TSet}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PolicyFilterSet= &CommandSpec{
	[]Token{TNo, TIpv6, TPolicy, TFilter, TSet}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

