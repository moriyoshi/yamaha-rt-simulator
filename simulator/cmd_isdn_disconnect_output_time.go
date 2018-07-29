package simulator

import "context"

var CmdIsdnDisconnectOutputTime = &CommandSpec{
	[]Token{TIsdn, TDisconnect, TOutput, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnDisconnectOutputTime = &CommandSpec{
	[]Token{TNo, TIsdn, TDisconnect, TOutput, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

