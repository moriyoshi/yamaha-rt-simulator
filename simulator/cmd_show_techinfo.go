package simulator

import "context"

var CmdShowTechinfo = &CommandSpec{
	[]Token{TShow, TTechinfo}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

