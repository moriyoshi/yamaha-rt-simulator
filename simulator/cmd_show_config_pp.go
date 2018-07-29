package simulator

import "context"

var CmdShowConfigPp = &CommandSpec{
	[]Token{TShow, TConfig, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdLessConfigPp = &CommandSpec{
	[]Token{TLess, TConfig, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
