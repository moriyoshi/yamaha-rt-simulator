package simulator

import "context"

var CmdSipSessionTimer = &CommandSpec{
	[]Token{TSip, TSession, TTimer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipSessionTimer = &CommandSpec{
	[]Token{TNo, TSip, TSession, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

