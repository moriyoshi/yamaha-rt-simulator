package simulator

import "context"

var CmdClearArp = &CommandSpec{
	[]Token{TClear, TArp}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

