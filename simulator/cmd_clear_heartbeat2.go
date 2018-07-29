package simulator

import "context"

var CmdClearHeartbeat2 = &CommandSpec{
	[]Token{TClear, THeartbeat2, &EitherToken{TId, TName}}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
