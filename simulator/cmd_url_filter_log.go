package simulator

import "context"

var CmdUrlFilterLog = &CommandSpec{
	[]Token{TUrl, TFilter, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUrlFilterLog = &CommandSpec{
	[]Token{TNo, TUrl, TFilter, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
