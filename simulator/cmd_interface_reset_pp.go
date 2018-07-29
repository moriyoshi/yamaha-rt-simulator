package simulator

import "context"

var CmdInterfaceResetPp = &CommandSpec{
	[]Token{TInterface, TReset, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

