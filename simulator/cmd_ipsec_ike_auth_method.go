package simulator

import "context"

var CmdIpsecIkeAuthMethod = &CommandSpec{
	[]Token{TIpsec, TIke, TAuth, TMethod}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeAuthMethod = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TAuth, TMethod}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

