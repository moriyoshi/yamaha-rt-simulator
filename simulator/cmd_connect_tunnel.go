package simulator

import "context"

var CmdConnectTunnel = &CommandSpec{
	[]Token{TConnect, TTunnel}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
