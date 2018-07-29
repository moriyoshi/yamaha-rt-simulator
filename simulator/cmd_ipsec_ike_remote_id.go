package simulator

import "context"

var CmdIpsecIkeRemoteId = &CommandSpec{
	[]Token{TIpsec, TIke, TRemote, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeRemoteId = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TRemote, TId}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

