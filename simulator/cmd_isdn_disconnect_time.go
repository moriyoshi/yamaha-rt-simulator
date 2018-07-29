package simulator

import "context"

var CmdIsdnDisconnectTime = &CommandSpec{
	[]Token{TIsdn, TDisconnect, TTime}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnDisconnectTime = &CommandSpec{
	[]Token{TNo, TIsdn, TDisconnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

