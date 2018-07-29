package simulator

import "context"

var CmdTunnelNgnRadiusAuth = &CommandSpec{
	[]Token{TTunnel, TNgn, TRadius, TAuth}, 4,
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

var CmdNoTunnelNgnRadiusAuth = &CommandSpec{
	[]Token{TNo, TTunnel, TNgn, TRadius, TAuth}, 5,
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
