package simulator

import "context"

var CmdAuthUserGroup = &CommandSpec{
	[]Token{TAuth, TUser, TGroup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAuthUserGroup = &CommandSpec{
	[]Token{TNo, TAuth, TUser, TGroup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

