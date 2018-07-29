package simulator

import "context"

var CmdShowAlias = &CommandSpec{
	[]Token{TShow, TAlias}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

