package simulator

import "context"

var CmdClearIpInboundFilter = &CommandSpec{
	[]Token{TClear, TIp, TInbound, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
