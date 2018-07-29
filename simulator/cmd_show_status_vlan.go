package simulator

import "context"

var CmdShowStatusVlan = &CommandSpec{
	[]Token{TShow, TStatus, TVlan}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

