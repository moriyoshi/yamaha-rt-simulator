package simulator

import "context"

var CmdIpsecIkeGroup = &CommandSpec{
	[]Token{TIpsec, TIke, TGroup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeGroup = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TGroup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

