package simulator

import "context"

var CmdSipServerRegisterRequestUri = &CommandSpec{
	[]Token{TSip, TServer, TRegister, TRequestUri}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerRegisterRequestUri = &CommandSpec{
	[]Token{TNo, TSip, TServer, TRegister, TRequestUri}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

