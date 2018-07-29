package simulator

import "context"

var CmdTelnetdListen = &CommandSpec{
	[]Token{TTelnetd, TListen}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoTelnetdListen = &CommandSpec{
	[]Token{TNo, TTelnetd, TListen}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

