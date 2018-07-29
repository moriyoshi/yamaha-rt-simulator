package simulator

import "context"

var CmdSsh = &CommandSpec{
	[]Token{TSsh}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
