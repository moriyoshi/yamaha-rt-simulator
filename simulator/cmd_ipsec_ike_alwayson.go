package simulator

import "context"

var CmdIpsecIkeAlwaysOn = &CommandSpec{
	[]Token{TIpsec, TIke, TAlwaysOn}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeAlwaysOn = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TAlwaysOn}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

