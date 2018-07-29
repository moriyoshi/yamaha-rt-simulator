package simulator

import "context"

var CmdIsdnForcedDisconnectTime = &CommandSpec{
	[]Token{TIsdn, TForced, TDisconnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnForcedDisconnectTime = &CommandSpec{
	[]Token{TNo, TIsdn, TForced, TDisconnect, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

