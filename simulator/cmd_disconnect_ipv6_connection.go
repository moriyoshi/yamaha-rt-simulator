package simulator

import "context"

var CmdDisconnectIpv6Connection = &CommandSpec{
	[]Token{TDisconnect, TIpv6, TConnection}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

