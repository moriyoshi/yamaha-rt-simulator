package simulator

import "context"

var CmdSetDefaultConfig = &CommandSpec{
	[]Token{TSetDefaultConfig}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

