package simulator

import "context"

var CmdShowStatusIpInboundFilter = &CommandSpec{
	[]Token{TShow, TStatus, TIp, TInbound, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusIpv6InboundFilter = &CommandSpec{
	[]Token{TShow, TStatus, TIpv6, TInbound, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

