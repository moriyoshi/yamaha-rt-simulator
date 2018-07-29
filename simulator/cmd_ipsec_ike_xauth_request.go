package simulator

import "context"

var CmdIpsecIkeXauthRequest = &CommandSpec{
	[]Token{TIpsec, TIke, TXauth, TRequest}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeXauthRequest = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TXauth, TRequest}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

