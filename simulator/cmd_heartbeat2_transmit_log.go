package simulator

import "context"

var CmdHeartbeat2TransmitLog = &CommandSpec{
	[]Token{THeartbeat2, TTransmit, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2TransmitLog = &CommandSpec{
	[]Token{TNo, THeartbeat2, TTransmit, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

