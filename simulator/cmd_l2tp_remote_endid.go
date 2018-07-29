package simulator

import "context"

var CmdL2TpRemoteEndId = &CommandSpec{
	[]Token{TL2Tp, TRemote, TEndId}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpRemoteEndId = &CommandSpec{
	[]Token{TNo, TL2Tp, TRemote, TEndId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

