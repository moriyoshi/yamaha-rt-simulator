package simulator

import "context"

var CmdShowPkiCrl = &CommandSpec{
	[]Token{TShow, TPki, TCrl}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

