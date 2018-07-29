package simulator

import "context"

var CmdTunnelNgnFallback = &CommandSpec{
	[]Token{TTunnel, TNgn, TFallback}, 3,
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

var CmdNoTunnelNgnFallback = &CommandSpec{
	[]Token{TNo, TTunnel, TNgn, TFallback}, 4,
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
