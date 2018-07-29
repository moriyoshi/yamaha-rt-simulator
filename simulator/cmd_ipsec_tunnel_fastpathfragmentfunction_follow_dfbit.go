package simulator

import "context"

var CmdIpsecTunnelFastpathFragmentFunctionFollowDfBit = &CommandSpec{
	[]Token{TIpsec, TTunnel, TFastpathFragmentFunction, TFollow, TDfBit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecTunnelFastpathFragmentFunctionFollowDfBit = &CommandSpec{
	[]Token{TNo, TIpsec, TTunnel, TFastpathFragmentFunction, TFollow, TDfBit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

