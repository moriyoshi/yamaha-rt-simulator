package simulator

import "context"

var CmdSshdService = &CommandSpec{
	[]Token{TSshd, TService}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshdService = &CommandSpec{
	[]Token{TNo, TSshd, TService}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

