package simulator

import "context"

var CmdClearInarp = &CommandSpec{
	[]Token{TClear, TInarp}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

