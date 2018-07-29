package simulator

import "context"

var CmdShowArp = &CommandSpec{
	[]Token{TShow, TArp}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

