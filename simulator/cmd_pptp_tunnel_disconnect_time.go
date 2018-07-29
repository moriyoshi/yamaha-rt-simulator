package simulator

import "context"

var CmdPptpTunnelDisconnectTime = &CommandSpec{
	[]Token{TPptp, TTunnel, TDisconnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpTunnelDisconnectTime = &CommandSpec{
	[]Token{TNo, TPptp, TTunnel, TDisconnect, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
