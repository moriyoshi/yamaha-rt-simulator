package simulator

import "context"

var CmdSave = &CommandSpec{
	[]Token{TSave}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

