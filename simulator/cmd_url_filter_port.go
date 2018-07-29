package simulator

import "context"

var CmdUrlFilterPort = &CommandSpec{
	[]Token{TUrl, TFilter, TPort}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUrlFilterPort = &CommandSpec{
	[]Token{TNo, TUrl, TFilter, TPort}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
