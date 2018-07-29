package simulator

import "context"

var CmdCopy = &CommandSpec{
	[]Token{TCopy}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

