package simulator

import "context"

var CmdPriLoopbackActive = &CommandSpec{
	[]Token{TPri, TLoopback, TActive}, 35,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
