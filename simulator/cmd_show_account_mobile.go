package simulator

import "context"

var CmdShowAccountMobile = &CommandSpec{
	[]Token{TShow, TAccount, TMobile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

