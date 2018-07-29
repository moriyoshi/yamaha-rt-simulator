package simulator

import "context"

var CmdIpsecIkeNegotiateStrictly = &CommandSpec{
	[]Token{TIpsec, TIke, TNegotiateStrictly}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeNegotiateStrictly = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TNegotiateStrictly}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

