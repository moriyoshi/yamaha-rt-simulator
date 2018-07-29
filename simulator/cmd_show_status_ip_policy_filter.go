package simulator

import "context"

var CmdShowStatusIpPolicyFilter = &CommandSpec{
	[]Token{TShow, TStatus, TIp, TPolicy, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusIpv6PolicyFilterIdType = &CommandSpec{
	[]Token{TShow, TStatus, TIpv6, TPolicy, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

