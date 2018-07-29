package simulator

import "context"

var CmdShowEnvironment = &CommandSpec{
	[]Token{TShow, TEnvironment}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

