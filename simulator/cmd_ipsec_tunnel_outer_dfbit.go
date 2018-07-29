package simulator

import "context"

var CmdIpsecTunnelOuterDfBit = &CommandSpec{
	[]Token{TIpsec, TTunnel, TOuter, TDfBit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecTunnelOuterDfBit = &CommandSpec{
	[]Token{TNo, TIpsec, TTunnel, TOuter, TDfBit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

