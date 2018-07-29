package simulator

import "context"

var CmdHeartbeat2ReceiveRecordLimit = &CommandSpec{
	[]Token{THeartbeat2, TReceive, TRecord, TLimit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2ReceiveRecordLimit = &CommandSpec{
	[]Token{TNo, THeartbeat2, TReceive, TRecord, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

