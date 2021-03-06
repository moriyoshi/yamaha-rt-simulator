package simulator

import "context"

var CmdShowHistory = &CommandSpec{
	[]Token{TShow, THistory}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

