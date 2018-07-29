package simulator

import "context"

var CmdHeartbeat2TransmitEnable = &CommandSpec{
	[]Token{THeartbeat2, TTransmit, TEnable}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2TransmitEnable = &CommandSpec{
	[]Token{TNo, THeartbeat2, TTransmit, TEnable}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

