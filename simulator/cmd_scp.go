package simulator

import "context"

var CmdScp = &CommandSpec{
	[]Token{TScp}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
