package simulator

import "context"

var CmdWanDisconnectInputTime = &CommandSpec{
	[]Token{TWan, TDisconnect, TInput, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanDisconnectInputTime = &CommandSpec{
	[]Token{TNo, TWan, TDisconnect, TInput, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

