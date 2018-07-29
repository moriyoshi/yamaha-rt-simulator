package simulator

import "context"

var CmdHeartbeatSend = &CommandSpec{
	[]Token{THeartbeat, TSend}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

