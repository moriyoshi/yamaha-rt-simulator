package simulator

import "context"

var CmdHeartbeat2ReceiveEnable = &CommandSpec{
	[]Token{THeartbeat2, TReceive, TEnable}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2ReceiveEnable = &CommandSpec{
	[]Token{TNo, THeartbeat2, TReceive, TEnable}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

