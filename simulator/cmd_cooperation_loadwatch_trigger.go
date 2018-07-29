package simulator

import "context"

var CmdCooperationLoadWatchTrigger = &CommandSpec{
	[]Token{TCooperation, TLoadWatch, TTrigger}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCooperationLoadWatchTrigger = &CommandSpec{
	[]Token{TNo, TCooperation, TLoadWatch, TTrigger}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

