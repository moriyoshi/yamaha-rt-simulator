package simulator

import "context"

var CmdSipUserAgent = &CommandSpec{
	[]Token{TSip, TUser, TAgent}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipUserAgent = &CommandSpec{
	[]Token{TNo, TSip, TUser, TAgent}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

