package simulator

import "context"

var CmdHeartbeat2Transmit = &CommandSpec{
	[]Token{THeartbeat2, TTransmit}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2Transmit = &CommandSpec{
	[]Token{TNo, THeartbeat2, TTransmit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

