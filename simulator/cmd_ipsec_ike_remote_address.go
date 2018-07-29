package simulator

import "context"

var CmdIpsecIkeRemoteAddress = &CommandSpec{
	[]Token{TIpsec, TIke, TRemote, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeRemoteAddress = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TRemote, TAddress}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

