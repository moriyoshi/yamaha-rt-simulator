package simulator

import "context"

var CmdIpPolicyFilterSetEnable = &CommandSpec{
	[]Token{TIp, TPolicy, TFilter, TSet, TEnable}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PolicyFilterSetEnable = &CommandSpec{
	[]Token{TIpv6, TPolicy, TFilter, TSet, TEnable}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPolicyFilterSetEnable = &CommandSpec{
	[]Token{TNo, TIp, TPolicy, TFilter, TSet, TEnable}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PolicyFilterSetEnable = &CommandSpec{
	[]Token{TNo, TIpv6, TPolicy, TFilter, TSet, TEnable}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

