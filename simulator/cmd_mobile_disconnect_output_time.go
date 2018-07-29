package simulator

import "context"

var CmdMobileDisconnectOutputTime = &CommandSpec{
	[]Token{TMobile, TDisconnect, TOutput, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileDisconnectOutputTime = &CommandSpec{
	[]Token{TNo, TMobile, TDisconnect, TOutput, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

