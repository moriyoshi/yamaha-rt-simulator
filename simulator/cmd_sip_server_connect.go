package simulator

import "context"

var CmdSipServerConnect = &CommandSpec{
	[]Token{TSip, TServer, TConnect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

