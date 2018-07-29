package simulator

import "context"

var CmdColdStart = &CommandSpec{
	[]Token{TCold, TStart}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

