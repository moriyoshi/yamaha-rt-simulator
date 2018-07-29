package simulator

import "context"

var CmdIpsecIkeVersion = &CommandSpec{
	[]Token{TIpsec, TIke, TVersion}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIKeVersion = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TVersion}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

