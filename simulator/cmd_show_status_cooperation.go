package simulator

import "context"

var CmdShowStatusCooperation = &CommandSpec{
	[]Token{TShow, TStatus, TCooperation}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
