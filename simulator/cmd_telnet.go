package simulator

import "context"

var CmdTelnet = &CommandSpec{
	[]Token{TTelnet}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

