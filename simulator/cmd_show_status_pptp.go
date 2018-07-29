package simulator

import "context"

var CmdShowStatusPptp = &CommandSpec{
	[]Token{TShow, TStatus, TPptp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

