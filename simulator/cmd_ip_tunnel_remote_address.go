package simulator

import "context"

var CmdIpTunnelRemoteAddress = &CommandSpec{
	[]Token{TIp, TTunnel, TRemote, TAddress}, 4,
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

var CmdNoIpTunnelRemoteAddress = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRemote, TAddress}, 5,
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
