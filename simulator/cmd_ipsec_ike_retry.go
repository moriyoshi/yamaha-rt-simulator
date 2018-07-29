package simulator

import "context"

var CmdIpsecIkeRetry = &CommandSpec{
	[]Token{TIpsec, TIke, TRetry}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeRetry = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TRetry}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

