package simulator

import "context"

var CmdAuthUserGroupAttribute = &CommandSpec{
	[]Token{TAuth, TUser, TGroup, TAttribute}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAuthUserGroupAttribute = &CommandSpec{
	[]Token{TNo, TAuth, TUser, TGroup, TAttribute}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

