package simulator

import "context"

var CmdShowStatusUpnp = &CommandSpec{
	[]Token{TShow, TStatus, TUpnp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

