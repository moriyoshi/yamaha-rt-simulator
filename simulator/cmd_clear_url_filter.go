package simulator

import "context"

var CmdClearUrlFilter = &CommandSpec{
	[]Token{TClear, TUrl, TFilter, &EitherToken{TPp, TTunnel}}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
