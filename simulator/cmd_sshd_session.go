package simulator

import "context"

var CmdSshdSession = &CommandSpec{
	[]Token{TSshd, TSession}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshdSession = &CommandSpec{
	[]Token{TNo, TSshd, TSession}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

