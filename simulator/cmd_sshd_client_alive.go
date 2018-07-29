package simulator

import "context"

var CmdSshdClientAlive = &CommandSpec{
	[]Token{TSshd, TClient, TAlive}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshdClientAlive = &CommandSpec{
	[]Token{TNo, TSshd, TClient, TAlive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

