package simulator

import "context"

var CmdIpsecIkeLocalId = &CommandSpec{
	[]Token{TIpsec, TIke, TLocal, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeLocalIdGateway_IdIp_AddressMask = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TLocal, TId}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

