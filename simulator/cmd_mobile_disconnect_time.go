package simulator

import "context"

var CmdMobileDisconnectTime = &CommandSpec{
	[]Token{TMobile, TDisconnect, TTime}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileDisconnectTime = &CommandSpec{
	[]Token{TNo, TMobile, TDisconnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

