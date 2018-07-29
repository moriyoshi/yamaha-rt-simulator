package simulator

import "context"

var CmdL2TpTunnelAuth = &CommandSpec{
	[]Token{TL2Tp, TTunnel, TAuth}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpTunnelAuth = &CommandSpec{
	[]Token{TNo, TL2Tp, TTunnel, TAuth}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
