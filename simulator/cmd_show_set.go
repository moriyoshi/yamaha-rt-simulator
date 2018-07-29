package simulator

import "context"

var CmdShowSet = &CommandSpec{
	[]Token{TShow, TSet}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

