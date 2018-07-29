package simulator

import "context"

var CmdShowStatusIpPolicyService = &CommandSpec{
	[]Token{TShow, TStatus, TIp, TPolicy, TService}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusIpv6PolicyService = &CommandSpec{
	[]Token{TShow, TStatus, TIpv6, TPolicy, TService}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

