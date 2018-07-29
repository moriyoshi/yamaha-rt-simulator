package simulator

import "context"

var CmdShowStatusMailService = &CommandSpec{
	[]Token{TShow, TStatus, TMail, TService}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

