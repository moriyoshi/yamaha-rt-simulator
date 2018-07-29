package simulator

import "context"

var CmdEcho = &CommandSpec{
	[]Token{TEcho}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

