package simulator

import "context"

var CmdHeartbeatReceive = &CommandSpec{
	[]Token{THeartbeat, TReceive}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeatReceive = &CommandSpec{
	[]Token{TNo, THeartbeat, TReceive}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

