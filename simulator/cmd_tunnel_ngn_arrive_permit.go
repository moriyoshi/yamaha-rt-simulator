package simulator

import "context"

var CmdTunnelNgnArrivePermit = &CommandSpec{
	[]Token{TTunnel, TNgn, TArrive, TPermit}, 4,
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

var CmdNoTunnelNgnArrivePermitPermit = &CommandSpec{
	[]Token{TNo, TTunnel, TNgn, TArrive, TPermit}, 4,
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
