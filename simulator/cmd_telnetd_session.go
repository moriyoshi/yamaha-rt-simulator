package simulator

import "context"

var CmdTelnetdSession = &CommandSpec{
	[]Token{TTelnetd, TSession}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoTelnetdSession = &CommandSpec{
	[]Token{TNo, TTelnetd, TSession}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

