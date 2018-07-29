package simulator

import "context"

var CmdConsoleInfo = &CommandSpec{
	[]Token{TConsole, TInfo}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoConsoleInfo = &CommandSpec{
	[]Token{TNo, TConsole, TInfo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

