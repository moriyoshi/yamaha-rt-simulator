package simulator

import "context"

var CmdDescriptionYno = &CommandSpec{
	[]Token{TDescription, TYno}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDescriptionYno = &CommandSpec{
	[]Token{TNo, TDescription, TYno}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

