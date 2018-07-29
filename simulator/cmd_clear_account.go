package simulator

import "context"

var CmdClearAccount = &CommandSpec{
	[]Token{TClear, TAccount}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
