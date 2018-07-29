package simulator

import "context"

var CmdDeleteExec = &CommandSpec{
	[]Token{TDelete, TExec}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

