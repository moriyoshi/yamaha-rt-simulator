package simulator

import "context"

var CmdIsdnPiafsControl = &CommandSpec{
	[]Token{TIsdn, TPiafs, TControl}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnPiafsControl = &CommandSpec{
	[]Token{TNo, TIsdn, TPiafs, TControl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

