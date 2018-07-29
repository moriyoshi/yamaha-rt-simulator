package simulator

import "context"

var CmdConfirm = &CommandSpec{
	[]Token{TConfirm}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

