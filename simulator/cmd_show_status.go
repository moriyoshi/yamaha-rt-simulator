package simulator

import "context"

var CmdShowStatus = &CommandSpec{
	[]Token{TShow, TStatus}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

