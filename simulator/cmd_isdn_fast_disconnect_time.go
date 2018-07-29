package simulator

import "context"

var CmdIsdnFastDisconnectTime = &CommandSpec{
	[]Token{TIsdn, TFast, TDisconnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNoIsdnFastDisconnectTime = &CommandSpec{
	[]Token{TNo, TNo, TIsdn, TFast, TDisconnect, TTime}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

