package simulator

import "context"

var CmdCooperationLoadWatchRemote = &CommandSpec{
	[]Token{TCooperation, TLoadWatch, TRemote}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCooperationLoadWatchRemote = &CommandSpec{
	[]Token{TNo, TCooperation, TLoadWatch, TRemote}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

