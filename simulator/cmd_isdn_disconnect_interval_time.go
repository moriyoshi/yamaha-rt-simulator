package simulator

import "context"

var CmdIsdnDisconnectIntervalTime = &CommandSpec{
	[]Token{TIsdn, TDisconnect, TInterval, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnDisconnectIntervalTime = &CommandSpec{
	[]Token{TNo, TIsdn, TDisconnect, TInterval, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

