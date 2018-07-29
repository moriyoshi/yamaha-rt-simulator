package simulator

import "context"

var CmdIsdnPiafsArrive = &CommandSpec{
	[]Token{TIsdn, TPiafs, TArrive}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnPiafsArrive = &CommandSpec{
	[]Token{TNo, TIsdn, TPiafs, TArrive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

