package simulator

import "context"

var CmdClearHeartbeat2Id = &CommandSpec{
	[]Token{TClear, THeartbeat2, TId}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
