package simulator

import "context"

var CmdDisconnectTunnel = &CommandSpec{
	[]Token{TDisconnect, TTunnel}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
