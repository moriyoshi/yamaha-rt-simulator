package simulator

import "context"

var CmdIsdnPiafsCall = &CommandSpec{
	[]Token{TIsdn, TPiafs, TCall}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnPiafsCall = &CommandSpec{
	[]Token{TNo, TIsdn, TPiafs, TCall}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

