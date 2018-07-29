package simulator

import "context"

var CmdShowStatusVrrp = &CommandSpec{
	[]Token{TShow, TStatus, TVrrp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

