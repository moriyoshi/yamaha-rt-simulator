package simulator

import "context"

var CmdShowStatusSd = &CommandSpec{
	[]Token{TShow, TStatus, TSd}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

