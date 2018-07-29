package simulator

import "context"

var CmdShowStatusBoot = &CommandSpec{
	[]Token{TShow, TStatus, TBoot}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

