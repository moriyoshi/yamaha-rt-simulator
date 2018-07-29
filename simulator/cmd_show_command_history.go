package simulator

import "context"

var CmdShowCommandHistory = &CommandSpec{
	[]Token{TShow, TCommand, THistory}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

