package simulator

import "context"

var CmdL2TpRemoteRouterId = &CommandSpec{
	[]Token{TL2Tp, TRemote, TRouterId}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpRemoteRouterId = &CommandSpec{
	[]Token{TNo, TL2Tp, TRemote, TRouterId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

