package simulator

import "context"

var CmdClearAccountMobile = &CommandSpec{
	[]Token{TClear, TAccount, TMobile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

