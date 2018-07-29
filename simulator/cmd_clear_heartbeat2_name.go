package simulator

import "context"

var CmdClearHeartbeat2Name = &CommandSpec{
	[]Token{TClear, THeartbeat2, TName}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
