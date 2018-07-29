package simulator

import "context"

var CmdRollbackTimer = &CommandSpec{
	[]Token{TRollback, TTimer}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
