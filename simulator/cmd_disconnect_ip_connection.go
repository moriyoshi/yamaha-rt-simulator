package simulator

import "context"

var CmdDisconnectIpConnection = &CommandSpec{
	[]Token{TDisconnect, TIp, TConnection}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

