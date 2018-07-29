package simulator

import "context"

var CmdSshdHideOpensshVersion = &CommandSpec{
	[]Token{TSshd, THide, TOpenssh, TVersion}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshdHideOpensshVersion = &CommandSpec{
	[]Token{TNo, TSshd, THide, TOpenssh, TVersion}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

