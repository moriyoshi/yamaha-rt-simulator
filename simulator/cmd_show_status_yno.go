package simulator

import "context"

var CmdShowStatusYno = &CommandSpec{
	[]Token{TShow, TStatus, TYno}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

