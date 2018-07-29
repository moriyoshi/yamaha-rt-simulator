package simulator

import "context"

var CmdUrlFilterReject = &CommandSpec{
	[]Token{TUrl, TFilter, TReject}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUrlFilterReject = &CommandSpec{
	[]Token{TNo, TUrl, TFilter, TReject}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
