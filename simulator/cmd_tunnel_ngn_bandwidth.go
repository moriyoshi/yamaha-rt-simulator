package simulator

import "context"

var CmdTunnelNgnBandwidth = &CommandSpec{
	[]Token{TTunnel, TNgn, TBandwidth}, 4,
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

var CmdNoTunnelNgnBandwidthBandwidth = &CommandSpec{
	[]Token{TNo, TTunnel, TNgn, TBandwidth}, 4,
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
