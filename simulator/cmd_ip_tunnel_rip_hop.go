package simulator

import "context"

var CmdIpTunnelRipHop = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, THop}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.Tunnel == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
