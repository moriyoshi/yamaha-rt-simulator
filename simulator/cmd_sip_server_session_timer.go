package simulator

import "context"

var CmdSipServerSessionTimer = &CommandSpec{
	[]Token{TSip, TServer, TSession, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerSessionTimer = &CommandSpec{
	[]Token{TNo, TSip, TServer, TSession, TTimer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

