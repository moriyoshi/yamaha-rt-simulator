package simulator

import "context"

var CmdSipServerDisplayName = &CommandSpec{
	[]Token{TSip, TServer, TDisplay, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerDisplayName = &CommandSpec{
	[]Token{TNo, TSip, TServer, TDisplay, TName}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

