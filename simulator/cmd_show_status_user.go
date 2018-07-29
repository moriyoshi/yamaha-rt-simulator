package simulator

import "context"

var CmdShowStatusUser = &CommandSpec{
	[]Token{TShow, TStatus, TUser}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

