package simulator

import "context"

var CmdShowLess = &CommandSpec{
	[]Token{TShow, TLess}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

