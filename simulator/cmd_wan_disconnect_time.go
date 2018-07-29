package simulator

import "context"

var CmdWanDisconnectTime = &CommandSpec{
	[]Token{TWan, TDisconnect, TTime}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanDisconnectTime = &CommandSpec{
	[]Token{TNo, TWan, TDisconnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
