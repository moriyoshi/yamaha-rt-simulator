package simulator

import "context"

var CmdHeartbeat2ReceiveLog = &CommandSpec{
	[]Token{THeartbeat2, TReceive, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2ReceiveLog = &CommandSpec{
	[]Token{TNo, THeartbeat2, TReceive, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

