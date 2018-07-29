package simulator

import "context"

var CmdDisconnectUser = &CommandSpec{
	[]Token{TDisconnect, TUser}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdDisconnectUserUserConnectionNo = &CommandSpec{
	[]Token{TDisconnect, TUser}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

