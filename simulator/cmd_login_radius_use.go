package simulator

import "context"

var CmdLoginRadiusUse = &CommandSpec{
	[]Token{TLogin, TRadius, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLoginRadiusUse = &CommandSpec{
	[]Token{TNo, TLogin, TRadius, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

