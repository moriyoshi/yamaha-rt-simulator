package simulator

import "context"

var CmdLoginPassword = &CommandSpec{
	[]Token{TLogin, TPassword}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

