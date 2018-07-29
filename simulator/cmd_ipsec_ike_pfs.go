package simulator

import "context"

var CmdIpsecIkePfs = &CommandSpec{
	[]Token{TIpsec, TIke, TPfs}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkePfs = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TPfs}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

