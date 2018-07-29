package simulator

import "context"

var CmdWanDisconnectOutputTime = &CommandSpec{
	[]Token{TWan, TDisconnect, TOutput, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanDisconnectOutputTime = &CommandSpec{
	[]Token{TNo, TWan, TDisconnect, TOutput, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
