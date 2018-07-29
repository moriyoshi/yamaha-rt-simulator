package simulator

import "context"

var CmdClearAccountTunnel = &CommandSpec{
	[]Token{TClear, TAccount, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

