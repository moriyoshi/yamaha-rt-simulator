package simulator

import "context"

var CmdIsdnAutoConnect = &CommandSpec{
	[]Token{TIsdn, TAuto, TConnect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnAutoConnect = &CommandSpec{
	[]Token{TNo, TIsdn, TAuto, TConnect}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

