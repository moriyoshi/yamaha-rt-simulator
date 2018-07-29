package simulator

import "context"

var CmdYnoAccessCode = &CommandSpec{
	[]Token{TYno, TAccess, TCode}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoYnoAccessCode = &CommandSpec{
	[]Token{TNo, TYno, TAccess, TCode}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
