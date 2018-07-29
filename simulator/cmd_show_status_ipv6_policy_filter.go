package simulator

import "context"

var CmdShowStatusIpv6PolicyFilter = &CommandSpec{
	[]Token{TShow, TStatus, TIpv6, TPolicy, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
