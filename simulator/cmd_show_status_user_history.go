package simulator

import "context"

var CmdShowStatusUserHistory = &CommandSpec{
	[]Token{TShow, TStatus, TUser, THistory}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

