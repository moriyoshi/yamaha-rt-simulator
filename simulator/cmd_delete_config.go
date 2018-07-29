package simulator

import "context"

var CmdDeleteConfig = &CommandSpec{
	[]Token{TDelete, TConfig}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

