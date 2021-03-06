package simulator

import "context"

var CmdHeartbeat2Receive = &CommandSpec{
	[]Token{THeartbeat2, TReceive}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2Receive = &CommandSpec{
	[]Token{TNo, THeartbeat2, TReceive}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

