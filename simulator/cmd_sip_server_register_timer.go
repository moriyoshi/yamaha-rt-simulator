package simulator

import "context"

var CmdSipServerRegisterTimer = &CommandSpec{
	[]Token{TSip, TServer, TRegister, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerRegisterTimer = &CommandSpec{
	[]Token{TNo, TSip, TServer, TRegister, TTimer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

