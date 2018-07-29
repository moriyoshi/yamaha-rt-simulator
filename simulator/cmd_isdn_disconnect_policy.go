package simulator

import "context"

var CmdIsdnDisconnectPolicy = &CommandSpec{
	[]Token{TIsdn, TDisconnect, TPolicy}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnDisconnectPolicy = &CommandSpec{
	[]Token{TNo, TIsdn, TDisconnect, TPolicy}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

