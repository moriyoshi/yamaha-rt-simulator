package simulator

import "context"

var CmdShowStatusPp = &CommandSpec{
	[]Token{TShow, TStatus, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

