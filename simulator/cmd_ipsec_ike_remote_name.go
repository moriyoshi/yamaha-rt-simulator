package simulator

import "context"

var CmdIpsecIkeRemoteName = &CommandSpec{
	[]Token{TIpsec, TIke, TRemote, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeRemoteName = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TRemote, TName}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

