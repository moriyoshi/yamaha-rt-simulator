package simulator

import "context"

var CmdIpsecIkeHash = &CommandSpec{
	[]Token{TIpsec, TIke, THash}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeHash = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, THash}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

