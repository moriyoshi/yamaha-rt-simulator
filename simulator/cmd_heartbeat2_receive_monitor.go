package simulator

import "context"

var CmdHeartbeat2ReceiveMonitor = &CommandSpec{
	[]Token{THeartbeat2, TReceive, TMonitor}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2ReceiveMonitor = &CommandSpec{
	[]Token{TNo, THeartbeat2, TReceive, TMonitor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
