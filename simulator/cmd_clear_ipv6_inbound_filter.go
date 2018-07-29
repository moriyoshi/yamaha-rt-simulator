package simulator

import "context"

var CmdClearIpv6InboundFilter = &CommandSpec{
	[]Token{TClear, TIpv6, TInbound, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
