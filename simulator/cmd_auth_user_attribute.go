package simulator

import "context"

var CmdAuthUserAttribute = &CommandSpec{
	[]Token{TAuth, TUser, TAttribute}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAuthUserAttribute = &CommandSpec{
	[]Token{TNo, TAuth, TUser, TAttribute}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

