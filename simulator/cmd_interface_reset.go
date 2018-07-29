package simulator

import "context"

var CmdInterfaceReset = &CommandSpec{
	[]Token{TInterface, TReset}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

