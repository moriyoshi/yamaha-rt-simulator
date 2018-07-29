package simulator

import "context"

var CmdUrlFilterUse = &CommandSpec{
	[]Token{TUrl, TFilter, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUrlFilterUse = &CommandSpec{
	[]Token{TNo, TUrl, TFilter, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
