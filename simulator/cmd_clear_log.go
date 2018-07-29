package simulator

import "context"

var CmdClearLog = &CommandSpec{
	[]Token{TClear, TLog}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

