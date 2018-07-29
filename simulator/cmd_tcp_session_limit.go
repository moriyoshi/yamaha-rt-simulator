package simulator

import "context"

var CmdTcpSessionLimit = &CommandSpec{
	[]Token{TTcp, TSession, TLimit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoTcpSessionLimit = &CommandSpec{
	[]Token{TNo, TTcp, TSession, TLimit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

