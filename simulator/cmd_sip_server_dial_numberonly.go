package simulator

import "context"

var CmdSipServerDialNumberOnly = &CommandSpec{
	[]Token{TSip, TServer, TDial, TNumberOnly}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerDialNumberOnly = &CommandSpec{
	[]Token{TNo, TSip, TServer, TDial, TNumberOnly}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

