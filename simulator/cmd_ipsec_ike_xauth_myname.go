package simulator

import "context"

var CmdIpsecIkeXauthMyname = &CommandSpec{
	[]Token{TIpsec, TIke, TXauth, TMyname}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeXauthMyname = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TXauth, TMyname}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

