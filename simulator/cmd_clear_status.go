package simulator

import "context"

var CmdClearStatus = &CommandSpec{
	[]Token{TClear, TStatus, &EitherToken{TPp, TTunnel}}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
