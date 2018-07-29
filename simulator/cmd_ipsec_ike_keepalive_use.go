package simulator

import "context"

var CmdIpsecIkeKeepaliveUse = &CommandSpec{
	[]Token{TIpsec, TIke, TKeepalive, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeKeepaliveUse = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TKeepalive, TUse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
