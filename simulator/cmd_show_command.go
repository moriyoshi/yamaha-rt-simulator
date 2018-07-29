package simulator

import "context"

var CmdShowCommand = &CommandSpec{
	[]Token{TShow, TCommand}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

