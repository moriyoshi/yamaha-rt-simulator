package simulator

import "context"

var CmdIsdnDsu = &CommandSpec{
	[]Token{TIsdn, TDsu}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnDsu = &CommandSpec{
	[]Token{TNo, TIsdn, TDsu}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

