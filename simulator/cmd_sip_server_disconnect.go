package simulator

import "context"

var CmdSipServerDisconnect = &CommandSpec{
	[]Token{TSip, TServer, TDisconnect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

