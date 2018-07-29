package simulator

import "context"

var CmdL2TpTunnelDisconnectTime = &CommandSpec{
	[]Token{TL2Tp, TTunnel, TDisconnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpTunnelDisconnectTimeTime = &CommandSpec{
	[]Token{TNo, TL2Tp, TTunnel, TDisconnect, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

