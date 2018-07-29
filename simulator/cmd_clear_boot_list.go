package simulator

import "context"

var CmdClearBootList = &CommandSpec{
	[]Token{TClear, TBoot, TList}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

