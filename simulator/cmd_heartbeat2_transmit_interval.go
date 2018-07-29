package simulator

import "context"

var CmdHeartbeat2TransmitInterval = &CommandSpec{
	[]Token{THeartbeat2, TTransmit, TInterval}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2TransmitInterval = &CommandSpec{
	[]Token{TNo, THeartbeat2, TTransmit, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
